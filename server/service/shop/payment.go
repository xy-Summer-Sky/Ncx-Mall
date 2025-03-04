package shop

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"hash/crc32"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shop"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WXPayService struct{}
// WechatPayRequest 微信支付请求数据结构（JSON格式）
type WechatPayRequest struct {
    AppID          string  `json:"appid"`
    MchID          string  `json:"mch_id"`
    NonceStr       string  `json:"nonce_str"`
    Sign           string  `json:"sign"`
    Body           string  `json:"body"`
    OutTradeNo     uint    `json:"out_trade_no"`
    TotalFee       float64 `json:"total_fee"` // 单位：分
    SpbillCreateIP string  `json:"spbill_create_ip"`
    NotifyURL      string  `json:"notify_url"`
    TradeType      string  `json:"trade_type"`
    UserId         uint    `json:"user_id"`
    Attach         string  `json:"attach"` // 附加字段，将用户ID等自定义数据放入
}

// WechatPayResponse 微信支付响应数据结构（JSON格式）
type WechatPayResponse struct {
    ReturnCode string `json:"return_code"`
    ReturnMsg  string `json:"return_msg"`
    AppID      string `json:"appid"`
    MchID      string `json:"mch_id"`
    NonceStr   string `json:"nonce_str"`
    Sign       string `json:"sign"`
    ResultCode string `json:"result_code"`
    PrepayID   string `json:"prepay_id"`
    TradeType  string `json:"trade_type"`
}

// WechatPayOrder 调用微信支付统一下单接口，发送 JSON 格式请求
func (s * WXPayService) WechatPayOrder(order shop.ShopOrder) (*WechatPayResponse, error) {
    // 构造请求结构体（假设订单 Price 单位为元，转换为分后传递）
    reqData := WechatPayRequest{
        AppID:          global.GVA_CONFIG.WXPay.AppID,
        MchID:          global.GVA_CONFIG.WXPay.MchID,
        NonceStr:       s.generateNonceStr(),
        Body:           "订单支付",
        OutTradeNo:     order.ID,
        TotalFee:       order.Price * 100,
        SpbillCreateIP: "127.0.0.1",
        NotifyURL:      global.GVA_CONFIG.WXPay.NotifyURL,
        TradeType:      "APP",
        UserId:         order.UserID,
        Attach:         strconv.Itoa(int(order.UserID)), // 将用户ID作为字符串放入附加字段
    }
    // 生成签名（示例：使用 crc32 模拟签名；实际应根据微信支付要求用 MD5 或 HMAC-SHA256）
    reqData.Sign = s.generateSign(reqData)

    jsonRequest, err := json.Marshal(reqData)
    if err != nil {
        return nil, err
    }

    client := &http.Client{
        Timeout:   15 * time.Second,
        Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
    }

    resp, err := client.Post("https://api.mch.weixin.qq.com/pay/unifiedorder", "application/json", bytes.NewReader(jsonRequest))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var wxResp WechatPayResponse
    if err := json.Unmarshal(body, &wxResp); err != nil {
        return nil, err
    }

    if wxResp.ReturnCode != "SUCCESS" || wxResp.ResultCode != "SUCCESS" {
        return &wxResp, errors.New(wxResp.ReturnMsg)
    }

    // 下单成功后可更新订单状态（最终支付结果仍以回调为准）
    if err := global.GVA_DB.Model(&shop.ShopOrder{}).
        Where("id = ?", order.ID).
        Update("Status", "prepaid").Error; err != nil {
        return &wxResp, err
    }

    return &wxResp, nil
}

// generateNonceStr 生成随机字符串
func (s * WXPayService) generateNonceStr() string {
    return strconv.FormatInt(time.Now().UnixNano(), 10)
}

// generateSign 模拟签名生成方法（实际需要按照微信支付要求处理）
func (s * WXPayService) generateSign(req WechatPayRequest) string {
    signStr := "appid=" + req.AppID +
        "&body=" + req.Body +
        "&mch_id=" + req.MchID +
        "&nonce_str=" + req.NonceStr +
        "&notify_url=" + req.NotifyURL +
        "&out_trade_no=" + strconv.FormatUint(uint64(req.OutTradeNo), 10) +
        "&spbill_create_ip=" + req.SpbillCreateIP +
        "&total_fee=" + strconv.Itoa(int(req.TotalFee)) +
        "&trade_type=" + req.TradeType +
        "&attach=" + req.Attach +
        "&key=" + global.GVA_CONFIG.WXPay.ApiV3Key

    checksum := crc32.ChecksumIEEE([]byte(signStr))
    return strconv.FormatUint(uint64(checksum), 16)
}

// ===========================================================
// 以下部分实现微信支付回调（v3版，JSON格式+加密资源）
// ===========================================================

type CallbackRequest struct {
    ID           string             `json:"id"`
    CreateTime   string             `json:"create_time"`
    ResourceType string             `json:"resource_type"`
    EventType    string             `json:"event_type"`
    Summary      string             `json:"summary"`
    Resource     EncryptedResource  `json:"resource"`
}

type EncryptedResource struct {
    Algorithm      string `json:"algorithm"`
    Ciphertext     string `json:"ciphertext"`
    AssociatedData string `json:"associated_data"`
    Nonce          string `json:"nonce"`
}

// PaymentResult 为解密后的业务数据结构
type PaymentResult struct {
    SpMchID         string `json:"sp_mchid"`
    SubMchID        string `json:"sub_mchid"`
    SpAppID         string `json:"sp_appid"`
    SubAppID        string `json:"sub_appid"`
    OutTradeNo      string `json:"out_trade_no"`
    TransactionID   string `json:"transaction_id"`
    TradeType       string `json:"trade_type"`
    TradeState      string `json:"trade_state"`
    TradeStateDesc  string `json:"trade_state_desc"`
    BankType        string `json:"bank_type"`
    Attach          string `json:"attach"`         // 附加字段，例如用户ID
    SuccessTime     string `json:"success_time"`   // 支付完成时间
}

// decryptResource 使用 AES-256-GCM 对微信回调中的密文进行解密
func (s * WXPayService) decryptResource(key, associatedData, nonce, ciphertext string) ([]byte, error) {
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
func (s * WXPayService) NotifyWeChatPay(c *gin.Context) {
    body, err := io.ReadAll(c.Request.Body)
    if err != nil {
        c.String(http.StatusBadRequest, "读取请求错误")
        return
    }
    var callback CallbackRequest
    if err := json.Unmarshal(body, &callback); err != nil {
        c.String(http.StatusBadRequest, "解析JSON错误")
        return
    }
    // 此处建议进行回调验签（跳过此部分示例）
    if callback.EventType != "TRANSACTION.SUCCESS" {
        c.String(http.StatusBadRequest, "事件类型不匹配")
        return
    }

    // 解密 resource 数据
    apiV3Key := global.GVA_CONFIG.WXPay.ApiV3Key // 确保配置中填入正确的APIv3密钥（32字节）
    plaintext, err := s.decryptResource(apiV3Key, callback.Resource.AssociatedData, callback.Resource.Nonce, callback.Resource.Ciphertext)
    if err != nil {
        c.String(http.StatusBadRequest, "解密失败: "+err.Error())
        return
    }

    var result PaymentResult
    if err := json.Unmarshal(plaintext, &result); err != nil {
        c.String(http.StatusBadRequest, "解析支付结果错误")
        return
    }

    // 根据解密后的 out_trade_no 更新订单状态
    orderID, err := strconv.Atoi(result.OutTradeNo)
    if err != nil {
        c.String(http.StatusBadRequest, "订单号转换错误")
        return
    }
    if err := global.GVA_DB.Model(&shop.ShopOrder{}).
        Where("id = ?", orderID).
        Update("Status", "paid").Error; err != nil {
        global.GVA_LOG.Error("更新订单状态失败", zap.Error(err))
        c.String(http.StatusInternalServerError, "更新订单状态失败")
        return
    }

    // 根据订单的支付金额更新对应用户的UserLevel
    var order shop.ShopOrder
    if err := global.GVA_DB.First(&order, orderID).Error; err == nil {
        // attach字段中传递的是用户ID
        userID, err := strconv.Atoi(result.Attach)
        if err == nil {
            var user system.SysUser
            if err := global.GVA_DB.First(&user, userID).Error; err == nil {
                newLevel := s.calculateUserLevel(order.Price)
                global.GVA_DB.Model(&user).Update("UserLevel", newLevel)
            }
        }
    }

    // 应答微信支付平台，通知回调已成功接收
    // 微信支付v3回调只要求返回200状态码；若验签失败或处理出错则返回非200状态码
    c.Status(http.StatusOK)
}

// calculateUserLevel 根据支付金额返回对应的用户等级
func (s * WXPayService) calculateUserLevel(price float64) int {
    if price >= 100 {
        return 3
    } else if price >= 50 {
        return 2
    }
    return 1
}