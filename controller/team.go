package controller

import (
	"github.com/dump-time/games-admin-server/log"
	"github.com/dump-time/games-admin-server/model"
	"github.com/dump-time/games-admin-server/services"
	"github.com/dump-time/games-admin-server/util"
	"github.com/gin-gonic/gin"
)

type CreateTeamReq struct {
	Name         string `json:"name"`
	Organization string `json:"organization"`
	Code         string `json:"code"`
}

const (
	createTeamErrorCode = 4401
)

func CreateTeamController(context *gin.Context) {
	var req CreateTeamReq
	if err := context.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		util.ParamsErrResp(context)
		return
	}

	if err := services.CreateTeam(&model.Team{
		Name:         req.Name,
		Organization: req.Organization,
		Code:         req.Code,
	}); err != nil {
		log.Error(err)
		util.FailedResp(context, createTeamErrorCode, "Create team error")
		return
	}

	util.SuccessResp(context, nil)
}
