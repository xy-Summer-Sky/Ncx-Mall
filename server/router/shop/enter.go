package shop

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	ShopRouter
}

var (
	shopApi = api.ApiGroupApp.ShopApiGroup.ShopOrderApi
)