package router

import (
	"github.com/dump-time/games-admin-server/controller"
	"github.com/gin-gonic/gin"
)

func initVolunteerRouter(apiGroup *gin.RouterGroup) *gin.RouterGroup {
	apiGroup.GET("/volunteers", controller.ListVolunteersController)

	volunteerRouter := apiGroup.Group("/volunteer")
	volunteerRouter.POST("/", controller.AddVolunteerController)
	volunteerRouter.DELETE("/:id", controller.DeleteVolunteerController)

	return volunteerRouter
}
