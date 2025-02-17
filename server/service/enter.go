package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/shop"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup				// 后台系统服务群
	ShopServiceGroup    shop.ServiceGroup				// 商城服务群
	TunnelService 		TunnelService					// 隧道服务
}
