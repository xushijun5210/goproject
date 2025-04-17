package router

import (
	"goproject/api"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitUserRouter(Router *gin.RouterGroup) {
	zap.S().Info("配置用户相关的url")
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("List", api.GetUserList)
		UserRouter.POST("Create", api.CreateUser)
		UserRouter.PUT("Update", api.UpdateUser)
		UserRouter.DELETE("Delete", api.DeleteUser)
	}

}
