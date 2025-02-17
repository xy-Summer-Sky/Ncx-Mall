package tunnel

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	TunnelApi
}

var (
	tunnelService = service.ServiceGroupApp.TunnelService
)