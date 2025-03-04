package shop

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ShopOrderApi
	
}

var (
	shopOrderService = service.ServiceGroupApp.ShopServiceGroup
	
)