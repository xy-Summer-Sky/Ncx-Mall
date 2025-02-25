package initialize

import (
	"net/http"
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/docs"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())

	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	systemRouter := router.RouterGroupApp.System
	tunnelRouter := router.RouterGroupApp.Tunnel
	shopRouter   := router.RouterGroupApp.Shop
	
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	// global.GVA_LOG.Info("use middleware cors")

	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)

	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())

	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}

	// 注册功能api路由

	{
		// systemRouter
		systemRouter.InitJwtRouter(PrivateGroup)                       // jwt相关路由
		systemRouter.InitUserRouter(PrivateGroup)                      // 注册用户路由
		systemRouter.InitCasbinRouter(PrivateGroup)                    // 权限相关路由

		// tunnelRouter
		tunnelRouter.InitTunnelRouter(PrivateGroup)

		// shopRouter
		shopRouter.InitShopRouter(PrivateGroup)
	}

	global.GVA_ROUTERS = Router.Routes()

	global.GVA_LOG.Info("router register success")
	return Router
}
