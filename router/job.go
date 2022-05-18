package router

import (
	"github.com/dump-time/games-admin-server/controller"
	"github.com/gin-gonic/gin"
)

func initJobRouter(apiGroup *gin.RouterGroup) {
	apiGroup.POST("/team/:teamID/job", controller.AddJob)
}
