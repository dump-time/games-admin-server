package middleware

import (
	"fmt"
	"strconv"

	"github.com/dump-time/games-admin-server/log"
	"github.com/dump-time/games-admin-server/util"
	"github.com/gin-gonic/gin"
)

func AuthCheck(context *gin.Context) {
	session := util.ContextSession(context)
	teamIDSession := session.Get("teamid")
	teamIDPath, err := strconv.Atoi(context.Param("teamID"))
	if err != nil {
		log.Error(err)
		util.ParamsErrResp(context)
		return
	}

	if teamIDSession == nil {
		log.Info("not login")
		util.NotLoginResp(context)
	} else if teamIDSession.(int64) != int64(teamIDPath) {
		if teamIDSession != int64(-1) {
			username := session.Get("user")
			log.Error(fmt.Sprintf("user %v doesn't has enough priviledge to access team %v", username, teamIDPath))
			util.NotAllowedResp(context)
		}
	}
}
