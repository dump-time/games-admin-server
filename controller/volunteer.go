package controller

import (
	"database/sql"
	"strconv"

	"github.com/dump-time/games-admin-server/log"
	"github.com/dump-time/games-admin-server/model"
	"github.com/dump-time/games-admin-server/services"
	"github.com/dump-time/games-admin-server/util"
	"github.com/gin-gonic/gin"
)

const (
	addVolunteerErrorCode    int = 4101
	listVolunteerErrorCode   int = 4102
	delVolunteerErrorCode    int = 4103
	updateVolunteerErrorCode int = 4104
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

type UpdateVolunteerReq struct {
	AddVolunteerReq

	TeamIDNew int `json:"team_id"`
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
		Intention:  req.Intention,
		Tel:        req.Tel,
		Experience: req.Experience,
		Avatar:     req.Avatar,
		Employment: req.Employment,
	}

	if teamID == -1 {
		volunteerData.TeamID = sql.NullInt64{Valid: false}
	} else {
		volunteerData.TeamID = sql.NullInt64{Int64: int64(teamID), Valid: true}
	}

	if err := services.AddVolunteer(&volunteerData); err != nil {
		log.Error(err)
		util.FailedResp(context, addVolunteerErrorCode, "Add volunteer error")
		return
	}

	util.SuccessResp(context, nil)
}

func ListVolunteersController(context *gin.Context) {
	// Extract data from request
	teamIDRaw := context.Param("teamID")
	teamID, err := strconv.Atoi(teamIDRaw)
	if err != nil {
		log.Error(err)
		util.ParamsErrResp(context)
		return
	}
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
	var nullableTeamID sql.NullInt64
	if teamID == -1 {
		nullableTeamID = sql.NullInt64{Valid: false}
	} else {
		nullableTeamID = sql.NullInt64{
			Int64: int64(teamID),
			Valid: true,
		}
	}

	// Get volunteers in databases
	volunteers, err := services.ListVolunteers(nullableTeamID, offset, pageSize)
	if err != nil {
		log.Error(err)
		util.FailedResp(context, listVolunteerErrorCode, "List volunteer error")
		return
	}

	var volunteerResp []gin.H
	for _, volunteer := range volunteers {
		volunteerResp = append(volunteerResp, gin.H{
			"id":         volunteer.ID,
			"name":       volunteer.Name,
			"gender":     volunteer.Gender,
			"intention":  volunteer.Intention,
			"job":        volunteer.JobID.Int64,
			"tel":        volunteer.Tel,
			"experience": volunteer.Experience,
			"team_id":    volunteer.TeamID.Int64,
			"avatar":     volunteer.Avatar,
			"id_number":  volunteer.IDNumber,
		})
	}

	util.SuccessResp(context, volunteerResp)
}

func DeleteVolunteerController(context *gin.Context) {
	teamIDRaw := context.Param("teamID")
	volunteerIDRaw := context.Param("id")

	teamID, err := strconv.Atoi(teamIDRaw)
	if err != nil {
		log.Error(err)
		util.ParamsErrResp(context)
		return
	}
	volunteerID, err := strconv.Atoi(volunteerIDRaw)
	if err != nil {
		log.Error(err)
		util.ParamsErrResp(context)
		return
	}

	var nullableTeamID sql.NullInt64
	if teamID == -1 {
		nullableTeamID.Valid = false
	} else {
		nullableTeamID.Int64 = int64(teamID)
		nullableTeamID.Valid = true
	}

	if err := services.DeleteVolunteer(nullableTeamID, uint(volunteerID)); err != nil {
		log.Error(err)
		util.FailedResp(context, delVolunteerErrorCode, "DeleteVolunteer error")
		return
	}

	util.SuccessResp(context, nil)
}

func UpdateVolunteerController(context *gin.Context) {
	// Extract data from request
	teamIDRaw := context.Param("teamID")
	volunteerIDRaw := context.Param("id")
	teamID, err := strconv.Atoi(teamIDRaw)
	if err != nil {
		log.Error(err)
		util.ParamsErrResp(context)
		return
	}
	volunteerID, err := strconv.Atoi(volunteerIDRaw)
	if err != nil {
		log.Error(err)
		util.ParamsErrResp(context)
		return
	}
	var nullableTeamID sql.NullInt64
	if teamID == -1 {
		nullableTeamID.Valid = false
	} else {
		nullableTeamID.Int64 = int64(teamID)
		nullableTeamID.Valid = true
	}
	var req UpdateVolunteerReq
	if err := context.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		util.ParamsErrResp(context)
		return
	}

	volunteer := model.Volunteer{
		Name:       req.Name,
		Gender:     req.Gender,
		Intention:  req.Intention,
		Tel:        req.Tel,
		Experience: req.Experience,
		Avatar:     req.Avatar,
		IDNumber:   req.IDNumber,
		Employment: req.Employment,
	}
	if req.TeamIDNew == -1 {
		volunteer.TeamID.Valid = false
	} else {
		volunteer.TeamID.Int64 = int64(req.TeamIDNew)
		volunteer.TeamID.Valid = true
	}
	if err := services.UpdateVolunteer(nullableTeamID, uint(volunteerID), &volunteer); err != nil {
		log.Error(err)
		util.FailedResp(context, updateVolunteerErrorCode, "Update volunteer error")
		return
	}

	util.SuccessResp(context, nil)
}
