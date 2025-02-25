package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/shop"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/tunnel"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	TunnelApiGroup  tunnel.ApiGroup
	ShopApiGroup 	shop.ApiGroup
}
