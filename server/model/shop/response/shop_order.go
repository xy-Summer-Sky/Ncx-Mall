package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/shop"

type CreateShopOrderResponse struct {
	ShopOrder shop.ShopOrder `json:"shoporder"`
}

type DeleteShopOrderResponse struct {}

type GetUserOrdersResponse struct {
	ShopOrders  []shop.ShopOrder 	`json:"shoporders"`
	Total 		int64 				`json:"total"`
}