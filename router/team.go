package router

import (
	"github.com/dump-time/games-admin-server/controller"
	"github.com/dump-time/games-admin-server/middleware"
	"github.com/gin-gonic/gin"
)

func initTeamRouter(apiGroup *gin.RouterGroup) *gin.RouterGroup {
	teamAPI := apiGroup.Group("/team/:teamID")
	{
		teamAPI.Use(middleware.AuthCheck)

		teamAPI.GET("/")    // Get team info
		teamAPI.PATCH("/")  // update team info
		teamAPI.DELETE("/") // Delte team
	}

	apiGroup.Use(middleware.CheckRootPriviledge)
	{
		apiGroup.GET("/teams", controller.ListTeamsController)  // List all teams
		apiGroup.POST("/team", controller.CreateTeamController) // Create team
	}

	return teamAPI
}
