package system

import (
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	{
		baseRouter.POST("register", baseApi.Register)
		baseRouter.POST("login", baseApi.Login)
		baseRouter.GET("captcha", baseApi.Captcha)
	}
	return baseRouter
}
