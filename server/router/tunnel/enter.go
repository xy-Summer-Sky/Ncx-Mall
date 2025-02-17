package tunnel

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	TunnelRouter
}

var (
	tunnalApi           	= api.ApiGroupApp.TunnelApiGroup
)