package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dump-time/games-admin-server/log"
	"github.com/dump-time/games-admin-server/model"
	"github.com/dump-time/games-admin-server/services"
	"github.com/dump-time/games-admin-server/util"
	"github.com/gin-gonic/gin"
)

type AddVolunteerReq struct {
	Name       string `json:"name"`
	Gender     bool   `json:"gender"`
	Intention  int    `json:"intention"`
	Tel        string `json:"tel"`
	Experience string `json:"experience"`
	Avatar     string `json:"avatar"`
	IDNumber   string `json:"id_number"`
	Employment string `json:"employment"`
}

func AddVolunteerController(context *gin.Context) {
	teamIDRaw := context.Param("teamID")

	teamID, err := strconv.Atoi(teamIDRaw)
	if err != nil {
		log.Error(err)
		util.ParamsErrResp(context)
		return
	}
	var req AddVolunteerReq
	if err := context.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		util.ParamsErrResp(context)
		return
	}

	// Extract data from request body
	volunteerData := model.Volunteer{
		Name:       req.Name,
		IDNumber:   req.IDNumber,
		Gender:     req.Gender,
		Employment: req.Employment,
		TeamID:     nil,
	}

	if err := services.AddVolunteer(&volunteerData); err != nil {
		log.Error(err)
		util.FailedResp(context, 4101, fmt.Sprint(err))
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"hello": teamID,
	})
}
