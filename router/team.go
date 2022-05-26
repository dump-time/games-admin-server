package router

import (
	"github.com/dump-time/games-admin-server/middleware"
	"github.com/gin-gonic/gin"
)

func initTeamRouter(apiGroup *gin.RouterGroup) *gin.RouterGroup {
	apiGroup.GET("/teams") // List all teams
	apiGroup.PUT("/team")  // Create team

	teamAPI := apiGroup.Group("/team/:teamID")
	{
		teamAPI.Use(middleware.AuthCheck)

		teamAPI.GET("/")    // Get team info
		teamAPI.PATCH("/")  // Get team info
		teamAPI.DELETE("/") // Get team info
	}

	return teamAPI
}
