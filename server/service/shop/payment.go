package shop

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	shopReq "github.com/flipped-aurora/gin-vue-admin/server/model/shop/request"
	shopRes "github.com/flipped-aurora/gin-vue-admin/server/model/shop/response"
	"go.uber.org/zap"
)

type WXPayService struct {
}

// WechatPayServerRequest 微信支付请求数据结构（JSON格式）

// WechatPayOrder 调用微信支付统一下单接口，发送 JSON 格式请求
func (s *WXPayService) WechatPayOrder(wechatPayServerRequest shopReq.WechatPayServerRequest) (*shopRes.WechatPayResponseURLCode, error) {
	client := &http.Client{
		Timeout:   15 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: false}},
	}
	orderJSON, err := json.Marshal(wechatPayServerRequest)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/v3/pay/transactions/native", bytes.NewReader(orderJSON))
	if err != nil {
		global.GVA_LOG.Error("创建请求失败", zap.Error(err))
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	//req.Header.Set("User-Agent", "YourUserAgent")

	resp, err := client.Do(req)
	if err != nil {
		global.GVA_LOG.Error("微信支付请求失败", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("读取响应体失败", zap.Error(err))
		return nil, err
	}

	global.GVA_LOG.Info("微信支付响应", zap.String("body", string(body)))

	var wxResp shopRes.WechatPayResponseURLCode
	if err := json.Unmarshal(body, &wxResp); err != nil {
		return nil, err
	}
	// 下单成功后可更新订单状态（最终支付结果仍以回调为准）
	if err := global.GVA_DB.Model(&shop.ShopOrder{}).
		Where("id = ?", wechatPayServerRequest.OutTradeNo).
		Update("status", 1).Error; err != nil {
		return &wxResp, err
	}

	return &wxResp, nil
}

// decryptResource 使用 AES-256-GCM 对微信回调中的密文进行解密
func (s *WXPayService) decryptResource(key, associatedData, nonce, ciphertext string) ([]byte, error) {
	keyBytes := []byte(key)
	ct, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	ad := []byte(associatedData)
	nce := []byte(nonce)
	plaintext, err := aead.Open(nil, nce, ct, ad)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

// NotifyWeChatPay 实现微信支付回调接口（v3版）
// 说明：微信支付回调会发送 JSON 格式的通知，其中 resource 部分经过AES-256-GCM加密。
// 本回调接口将解密业务数据，从中获取订单号及附加数据（attach），更新订单状态并调整用户等级。
func (s *WXPayService) NotifyWeChatPay(c shopRes.WeChatNotifyResponse) {
	// Decrypt the resource part
	plaintext, err := s.decryptResource(
		global.GVA_CONFIG.WXPay.ApiV3Key,
		c.Resource.AssociatedData,
		c.Resource.Nonce,
		c.Resource.Ciphertext,
	)
	if err != nil {
		global.GVA_LOG.Error("解密资源失败", zap.Error(err))
		return
	}

	// Unmarshal the decrypted data into WeChatNotifyResponseDataResource
	var dataResource shopRes.WeChatNotifyResponseDataResource
	if err := json.Unmarshal(plaintext, &dataResource); err != nil {
		global.GVA_LOG.Error("解析解密数据失败", zap.Error(err))
		return
	}

	if err := global.GVA_DB.Model(&shop.ShopOrder{}).
		Where("id = ?", dataResource.OutTradeNo).
		Update("Status", dataResource.TradeState).Error; err != nil {
		global.GVA_LOG.Error("更新订单状态失败", zap.Error(err))
		return
	}
	// Convert attach field to uint
	userID, err := strconv.ParseUint(dataResource.Attach, 10, 32)
	if err != nil {
		global.GVA_LOG.Error("解析用户ID失败", zap.Error(err))
		return
	}
	// Adjust user level based on the payment amount (if needed)
	userLevel := s.calculateUserLevel(float64(dataResource.Amount.Total) / 100)
	if err := global.GVA_DB.Model(&system.SysUser{}).
		Where("uuid = ?", userID).
		Update("accountStatus", userLevel).Error; err != nil {
		global.GVA_LOG.Error("更新用户等级失败", zap.Error(err))
	}

}

// calculateUserLevel 根据支付金额返回对应的用户等级
func (s *WXPayService) calculateUserLevel(price float64) int {
	if price >= 100 {
		return 3
	} else if price >= 50 {
		return 2
	}
	return 1
}
