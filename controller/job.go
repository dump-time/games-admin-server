package controller

import (
	"database/sql"
	"github.com/dump-time/games-admin-server/model"
	"github.com/dump-time/games-admin-server/services"
	"github.com/dump-time/games-admin-server/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type addRequest struct {
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

	var req addRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
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
			TeamID:   sql.NullInt64{Int64: int64(teamId), Valid: true},
			Name:     req.Name,
			Content:  req.Content,
			Location: req.Location,
		}
	}

	if err := services.AddJob(mod); err != nil {
		util.FailedResp(ctx, 4201, "Add Job Failed")
		return
	}

	util.SuccessResp(ctx, nil)
}

func GetJobs(ctx *gin.Context) {
	param := ctx.Param("teamID")
	teamId, err := strconv.Atoi(param)
	if err != nil {
		util.ParamsErrResp(ctx)
		return
	}

	jobs, err := services.GetJobs(sql.NullInt64{
		Int64: int64(teamId),
		Valid: teamId >= 0,
	})
	if err != nil {
		util.FailedResp(ctx, 4202, "Get Jobs Failed")
		return
	}
	if len(jobs) == 0 {
		util.NotFoundResp(ctx)
		return
	}

	util.SuccessResp(ctx, jobs)
}
