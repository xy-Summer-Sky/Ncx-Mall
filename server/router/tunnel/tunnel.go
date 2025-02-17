package tunnel

import "github.com/gin-gonic/gin"

type TunnelRouter struct {}

func (s *TunnelRouter) InitTunnelRouter(Router *gin.RouterGroup) {
	tunnelRouter := Router.Group("tunnel")
	{
		tunnelRouter.POST("createTunnel", tunnalApi.CreateTunnel)
		tunnelRouter.POST("deleteTunnel", tunnalApi.DeleteTunnel)
		tunnelRouter.GET("findUserAllTunnel", tunnalApi.FindUserAllTunnels)
	}	
}