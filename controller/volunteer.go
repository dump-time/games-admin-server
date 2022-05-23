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
	searchVolunteerErrorCode int = 4105
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
	Status     int    `json:"status"`
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
		Name:     req.Name,
		IDNumber: req.IDNumber,
		Gender:   req.Gender,
		// IntentionID: req.Intention,
		Tel:        req.Tel,
		Experience: req.Experience,
		Avatar:     req.Avatar,
		Employment: req.Employment,
		Status:     req.Status,
	}
	if req.Intention == -1 {
		volunteerData.IntentionID.Valid = false
	} else {
		volunteerData.IntentionID.Valid = true
		volunteerData.IntentionID.Int64 = int64(req.Intention)
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

	// Get volunteer pages
	pagesNum, err := services.GetVolunteersNum(nullableTeamID, pageSize)
	if err != nil {
		log.Error(err)
		util.FailedResp(context, listVolunteerErrorCode, "List volunteer error")
		return
	}

	var volunteerList []gin.H
	for _, volunteer := range volunteers {
		volunteerData := gin.H{
			"id":         volunteer.ID,
			"name":       volunteer.Name,
			"gender":     volunteer.Gender,
			"intention":  volunteer.Intention.Name,
			"job":        volunteer.Job.Name,
			"tel":        volunteer.Tel,
			"experience": volunteer.Experience,
			"avatar":     volunteer.Avatar,
			"id_number":  volunteer.IDNumber,
			"status":     volunteer.Status,
			"employment": volunteer.Employment,
		}
		if volunteer.TeamID.Valid {
			volunteerData["team_id"] = volunteer.TeamID.Int64
		} else {
			volunteerData["team_id"] = nil
		}
		volunteerList = append(volunteerList, volunteerData)
	}

	util.SuccessResp(context, gin.H{
		"num":        pagesNum,
		"volunteers": volunteerList,
	})
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
		Tel:        req.Tel,
		Experience: req.Experience,
		Avatar:     req.Avatar,
		IDNumber:   req.IDNumber,
		Employment: req.Employment,
		Status:     req.Status,
	}
	if req.TeamIDNew == -1 {
		volunteer.TeamID.Valid = false
	} else {
		volunteer.TeamID.Int64 = int64(req.TeamIDNew)
		volunteer.TeamID.Valid = true
	}
	if req.Intention == -1 {
		volunteer.IntentionID.Valid = false
	} else {
		volunteer.TeamID.Int64 = int64(req.Intention)
		volunteer.TeamID.Valid = true
	}

	if err := services.UpdateVolunteer(nullableTeamID, uint(volunteerID), &volunteer); err != nil {
		log.Error(err)
		util.FailedResp(context, updateVolunteerErrorCode, "Update volunteer error")
		return
	}

	util.SuccessResp(context, nil)
}

func SearchVolunteerController(context *gin.Context) {
	// Extract data from request
	teamIDRaw := context.Param("teamID")
	IDNumber := context.Param("IDNumber")
	teamID, err := strconv.Atoi(teamIDRaw)
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

	volunteer, error := services.SearchVolunteer(nullableTeamID, IDNumber)
	if error != nil {
		log.Error(err)
		util.FailedResp(context, searchVolunteerErrorCode, "Update volunteer error")
		return
	}

	resp := gin.H{
		"id":         volunteer.ID,
		"name":       volunteer.Name,
		"gender":     volunteer.Gender,
		"intention":  volunteer.IntentionID,
		"tel":        volunteer.Tel,
		"experience": volunteer.Experience,
		"avatar":     volunteer.Avatar,
		"id_number":  volunteer.IDNumber,
		"employment": volunteer.Employment,
		"status":     volunteer.Status,
	}
	if volunteer.TeamID.Valid {
		resp["team_id"] = volunteer.TeamID.Int64
	} else {
		resp["team_id"] = nil
	}
	util.SuccessResp(context, resp)
}
