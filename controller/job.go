package controller

import (
	"database/sql"
	"errors"
	"github.com/dump-time/games-admin-server/log"
	"github.com/dump-time/games-admin-server/model"
	"github.com/dump-time/games-admin-server/services"
	"github.com/dump-time/games-admin-server/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

func getTeamID(ctx *gin.Context) (int64, error) {
	param := ctx.Param("teamID")
	teamId, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}
	return int64(teamId), nil
}

type addRequest struct {
	Name     string `json:"name"`
	Content  string `json:"content"`
	Location string `json:"location"`
}

func AddJob(ctx *gin.Context) {
	teamId, err := getTeamID(ctx)
	if err != nil {
		_ = ctx.Error(err)
		util.ParamsErrResp(ctx)
		return
	}

	var req addRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		util.ParamsErrResp(ctx)
		return
	}

	var mod *model.Job

	if teamId < 0 {
		mod = &model.Job{
			Name:     req.Name,
			Content:  req.Content,
			Location: req.Location,
		}
	} else {
		mod = &model.Job{
			TeamID:   sql.NullInt64{Int64: teamId, Valid: true},
			Name:     req.Name,
			Content:  req.Content,
			Location: req.Location,
		}
	}

	if err := services.AddJob(mod); err != nil {
		_ = ctx.Error(err)
		util.FailedResp(ctx, 4201, "Add Job Failed")
		return
	}

	util.SuccessResp(ctx, nil)
}

func GetJobs(ctx *gin.Context) {
	teamId, err := getTeamID(ctx)
	if err != nil {
		_ = ctx.Error(err)
		util.ParamsErrResp(ctx)
		return
	}

	offsetRaw := ctx.DefaultQuery("offset", "0")
	pageSizeRaw := ctx.DefaultQuery("page-size", "10")
	offset, err := strconv.Atoi(offsetRaw)
	if err != nil {
		_ = ctx.Error(err)
		util.ParamsErrResp(ctx)
		return
	}
	pageSize, err := strconv.Atoi(pageSizeRaw)
	if err != nil {
		_ = ctx.Error(err)
		util.ParamsErrResp(ctx)
		return
	}

	// Super Admin
	if teamId == -1 {
		jobs, err := services.GetAllJobs(offset, pageSize)
		if err != nil {
			_ = ctx.Error(err)
			util.FailedResp(ctx, 4202, "Get Jobs Failed")
			return
		}
		util.SuccessResp(ctx, gin.H{
			"num":  len(jobs),
			"jobs": jobs,
		})
		return
	}

	jobs, err := services.GetJobs(sql.NullInt64{
		Int64: teamId,
		Valid: teamId >= 0,
	}, offset, pageSize)
	if err != nil {
		_ = ctx.Error(err)
		util.FailedResp(ctx, 4202, "Get Jobs Failed")
		return
	}
	if len(jobs) == 0 {
		util.SuccessResp(ctx, nil)
		return
	}

	num, err := services.GetJobsNum(sql.NullInt64{
		Int64: teamId,
		Valid: teamId >= 0,
	})
	if err != nil {
		_ = ctx.Error(err)
		util.FailedResp(ctx, 4202, "Get Jobs num Failed")
		return
	}

	util.SuccessResp(ctx, gin.H{
		"num":  num,
		"jobs": jobs,
	})
}

func getID(ctx *gin.Context) (uint, error) {
	param := ctx.Param("id")
	jobId, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}
	if jobId <= 0 {
		return 0, errors.New("invalid ID Error")
	}
	return uint(jobId), nil
}

func DeleteJob(ctx *gin.Context) {
	teamId, err := getTeamID(ctx)
	if err != nil {
		_ = ctx.Error(err)
		util.ParamsErrResp(ctx)
		return
	}

	id, err := getID(ctx)
	if err != nil {
		_ = ctx.Error(err)
		util.ParamsErrResp(ctx)
		return
	}

	rows, err := services.DeleteJob(sql.NullInt64{Int64: teamId, Valid: teamId >= 0}, id)
	if err != nil {
		_ = ctx.Error(err)
		util.FailedResp(ctx, 4203, "Delete Job Failed")
		return
	}
	if rows == 0 {
		util.SuccessResp(ctx, nil)
		return
	}

	util.SuccessResp(ctx, nil)
}

type updateRequest struct {
	Name     string `json:"name"`
	Content  string `json:"content"`
	Location string `json:"location"`
	TeamID   int64  `json:"team_id"`
}

func UpdateJob(ctx *gin.Context) {
	teamId, err := getTeamID(ctx)
	if err != nil {
		_ = ctx.Error(err)
		util.ParamsErrResp(ctx)
		return
	}

	id, err := getID(ctx)
	if err != nil {
		_ = ctx.Error(err)
		util.ParamsErrResp(ctx)
		return
	}

	var req updateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(err)
		util.ParamsErrResp(ctx)
		return
	}

	var mod *model.Job
	if teamId == -1 {
		mod = &model.Job{
			Name:     req.Name,
			Content:  req.Content,
			Location: req.Location,
			TeamID: sql.NullInt64{Int64: req.TeamID,
				Valid: req.TeamID >= 0},
		}
	} else {
		mod = &model.Job{
			Name:     req.Name,
			Content:  req.Content,
			Location: req.Location,
		}
	}

	log.Info(mod)

	rows, err := services.UpdateJob(
		id,
		mod)
	if err != nil {
		_ = ctx.Error(err)
		util.ParamsErrResp(ctx)
		return
	} else if rows == 0 {
		util.SuccessResp(ctx, nil)
		return
	}

	util.SuccessResp(ctx, nil)
}
