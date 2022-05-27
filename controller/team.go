package controller

import (
	"strconv"

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

func ListTeamsController(context *gin.Context) {
	offsetRaw := context.DefaultQuery("offset", "0")
	pageSizeRaw := context.DefaultQuery("page-size", "10")
	offset, err := strconv.Atoi(offsetRaw)
	if err != nil {
		log.Error(err)
		util.ParamsErrResp(context)
		return
	}
	pageSize, err := strconv.Atoi(pageSizeRaw)
	if err != nil {
		log.Error(err)
		util.ParamsErrResp(context)
		return
	}

	teams, num, err := services.ListTeams(offset, pageSize)
	if err != nil {
		log.Error(err)
		util.ParamsErrResp(context)
		return
	}

	var teamData []gin.H
	for _, team := range teams {
		teamData = append(teamData, gin.H{
			"name":         team.Name,
			"organization": team.Organization,
			"code":         team.Code,
		})
	}

	util.SuccessResp(context, gin.H{
		"num":   num,
		"teams": teamData,
	})
}
