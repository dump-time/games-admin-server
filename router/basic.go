package router

import (
	"github.com/dump-time/games-admin-server/controller"
	"github.com/gin-gonic/gin"
)

func initBasicRouter(apiGroup *gin.RouterGroup) *gin.RouterGroup {
	apiGroup.POST("/login", controller.LoginController)
	apiGroup.GET("/info", controller.GetAdminInfo)
	return apiGroup
}
