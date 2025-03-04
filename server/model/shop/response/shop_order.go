package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/shop"

type CreateShopOrderResponse struct {
	ShopOrder shop.ShopOrder `json:"shoporder"`
}

type DeleteShopOrderResponse struct{}

type GetUserOrdersResponse struct {
	ShopOrders []shop.ShopOrder `json:"shoporders"`
	Total      int64            `json:"total"`
}

// WechatPayResponseURLCode 微信支付响应数据结构（JSON格式）
type WechatPayResponseURLCode struct {
	CodeURL string `json:"code_url" `
}

type WeChatNotifyResponse struct {
	ID           string `json:"id"`
	CreateTime   string `json:"create_time"`
	ResourceType string `json:"resource_type"`
	EventType    string `json:"event_type"`
	Summary      string `json:"summary"`
	Resource     struct {
		OriginalType   string `json:"original_type"`
		Algorithm      string `json:"algorithm"`
		Ciphertext     string `json:"ciphertext"`
		AssociatedData string `json:"associated_data"`
		Nonce          string `json:"nonce"`
	} `json:"resource"`
}

type WeChatNotifyResponseDataResource struct {
	AppID          string `json:"appid" binding:"required"`            // 应用ID
	MchID          string `json:"mchid" binding:"required"`            // 商户号
	OutTradeNo     string `json:"out_trade_no" binding:"required"`     // 商户订单号
	TransactionID  string `json:"transaction_id" binding:"required"`   // 微信支付订单号
	TradeType      string `json:"trade_type" binding:"required"`       // 交易类型
	TradeState     string `json:"trade_state" binding:"required"`      // 交易状态
	TradeStateDesc string `json:"trade_state_desc" binding:"required"` // 交易状态描述
	BankType       string `json:"bank_type" binding:"required"`        // 付款银行
	Attach         string `json:"attach"`                              // 附加数据
	SuccessTime    string `json:"success_time" binding:"required"`     // 支付完成时间
	Payer          struct {
		OpenID string `json:"openid" binding:"required"`
	} `json:"payer" binding:"required"` // 支付者信息
	Amount struct {
		Total         int    `json:"total" binding:"required"`          // 订单总金额，单位为分
		PayerTotal    int    `json:"payer_total" binding:"required"`    // 用户支付金额，单位为分
		Currency      string `json:"currency" binding:"required"`       // 货币类型，CNY：人民币
		PayerCurrency string `json:"payer_currency" binding:"required"` // 用户支付币种
	} `json:"amount" binding:"required"` // 订单金额信息

}
