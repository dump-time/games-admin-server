package router

import (
	"github.com/dump-time/games-admin-server/middleware"
	"github.com/gin-gonic/gin"
)

func initTeamRouter(apiGroup *gin.RouterGroup) *gin.RouterGroup {
	teamAPI := apiGroup.Group("/team/:teamID")
	{
		teamAPI.Use(middleware.AuthCheck)

		teamAPI.GET("/")    // Get team info
		teamAPI.PATCH("/")  // Get team info
		teamAPI.DELETE("/") // Get team info
	}

	apiGroup.Use(middleware.CheckRootPriviledge)
	{
		apiGroup.GET("/teams") // List all teams
		apiGroup.POST("/team") // Create team
	}

	return teamAPI
}
