package controller

import (
	"net/http"
	"strconv"

	"github.com/dump-time/games-admin-server/util"
	"github.com/gin-gonic/gin"
)

func AddVolunteerController(context *gin.Context) {
	teamIDRaw := context.Param("teamID")

	teamID, err := strconv.Atoi(teamIDRaw)
	if err != nil {
		util.ParamsErrResp(context)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"hello": teamID,
	})
}
