package initialize

import (
	"goproject/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	Router := gin.Default()
	Router.SetTrustedProxies([]string{"127.0.0.1"})
	ApiGroup := Router.Group("/v1")
	router.InitUserRouter(ApiGroup)
	return Router
}
