package router

import (
	"github.com/dump-time/games-admin-server/controller"
	"github.com/gin-gonic/gin"
)

func initJobRouter(apiGroup *gin.RouterGroup) {
	apiGroup.POST("/job", controller.AddJob)
	apiGroup.GET("/jobs", controller.GetJobs)
}
