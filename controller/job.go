package controller

import (
	"github.com/dump-time/games-admin-server/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type request struct {
	Name     string `json:"name"`
	Content  string `json:"content"`
	Location string `json:"location"`
}

func AddJob(ctx *gin.Context) {
	param := ctx.Param("teamID")
	teamId, err := strconv.Atoi(param)
	if err != nil {
		util.ParamsErrResp(ctx)
		return
	}

	var req request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.ParamsErrResp(ctx)
		return
	}

	util.SuccessResp(ctx, gin.H{"team": teamId, "req": req})
}
