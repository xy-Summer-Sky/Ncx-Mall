package system

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	JwtApi
	BaseApi
	CasbinApi
}

var (
	jwtService              = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userService             = service.ServiceGroupApp.SystemServiceGroup.UserService
	casbinService           = service.ServiceGroupApp.SystemServiceGroup.CasbinService
)
