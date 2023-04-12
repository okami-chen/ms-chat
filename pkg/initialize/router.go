package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/okamin-chen/chat/app/controller"
	"github.com/okamin-chen/chat/app/middleware"
	"github.com/okamin-chen/chat/pkg/global"
)

func InitRouter() *gin.Engine {

	gin.ForceConsoleColor()
	gin.SetMode(global.Conf.Server.Model)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.ExampleMiddleware, middleware.LoggerMiddleware)

	UserRouterGroupInit(router)
	DefaultRouterGroupInit(router)
	AuthRouterGroupInit(router)

	return router
}

func DefaultRouterGroupInit(router *gin.Engine) {
	group := router.Group("/")
	{
		group.GET("/ping", controller.DefaultController{}.Ping)
	}
}

func UserRouterGroupInit(router *gin.Engine) {
	group := router.Group("/user")
	{
		group.GET("/info", controller.UserController{}.GetUserInfo)
	}
}

func AuthRouterGroupInit(router *gin.Engine) {
	group := router.Group("/auth")
	{
		group.GET("/wechat", controller.AuthController{}.LoginByWechat)
	}
}
