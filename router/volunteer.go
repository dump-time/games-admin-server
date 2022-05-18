package router

import (
	"github.com/dump-time/games-admin-server/controller"
	"github.com/gin-gonic/gin"
)

func initVolunteerRouter(apiGroup *gin.RouterGroup) {
	apiGroup.POST("/team/:teamID/volunteer", controller.AddVolunteerController)
}
