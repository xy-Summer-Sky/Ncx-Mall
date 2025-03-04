package request

// CreateOrder 创建订单请求结构
type CreateOrder struct {
	ServiceType uint    `json:"serviceType" example:"1" binding:"required"` // 服务类型：1-包年 2-包月
	Price       float64 `json:"price" example:"99.99" binding:"required"`   // 订单价格
	Status      uint    `json:"status" example:"0"`                         // 订单状态：0-待支付 1-已支付 2-已取消
	CreateTime  int64   `json:"createTime" example:"1645084800"`            // 订单创建时间
	ExpireTime  int64   `json:"expireTime" example:"1676620800"`            // 订单过期时间
}

// DeleteOrder 删除订单请求结构
type DeleteOrder struct {
	OrderID uint `json:"orderId" binding:"required"` // 订单ID
}

// GetUserOrders 获取用户订单请求结构
type GetUserOrders struct {
	UserID uint `json:"userId" binding:"required"` // 用户ID
	// Page   int  `json:"page" example:"1"`         // 页码
	// Size   int  `json:"size" example:"10"`        // 每页数量
}

type WechatPayWebRequest struct {
	Description string `json:"description" binding:"required"`  // 商品描述
	OutTradeNo  string `json:"out_trade_no" binding:"required"` // 商户订单号
	TimeExpire  string `json:"time_expire"`                     // 交易结束时间
	Attach      string `json:"attach"`                          // 附加数据
	Amount      struct {
		Total    int    `json:"total" binding:"required"`    // 总金额
		Currency string `json:"currency" binding:"required"` // 货币类型
	} `json:"amount" binding:"required"`
}

type WechatPayServerRequest struct {
	AppID     string `json:"appid" binding:"required"`      // 应用ID
	MchID     string `json:"mchid" binding:"required"`      // 直连商户号
	NotifyURL string `json:"notify_url" binding:"required"` // 通知地址
	WechatPayWebRequest
}
