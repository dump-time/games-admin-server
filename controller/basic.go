package controller

import (
	"github.com/dump-time/games-admin-server/log"
	"github.com/dump-time/games-admin-server/services"
	"github.com/dump-time/games-admin-server/util"
	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

const (
	loginErrorCode int = 4301
)

func LoginController(context *gin.Context) {
	var req LoginReq
	if err := context.ShouldBindJSON(&req); err != nil {
		log.Error(err)
		util.ParamsErrResp(context)
		return
	}

	err := services.CheckAuth(context, req.Username, req.Password)
	if err != nil {
		log.Error(err)
		util.FailedResp(context, loginErrorCode, "Login error")
		return
	}

	util.SuccessResp(context, nil)
}
