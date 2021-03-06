package services

import (
	"errors"

	"github.com/dump-time/games-admin-server/global"
	"github.com/dump-time/games-admin-server/model"
	"github.com/dump-time/games-admin-server/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CheckAuth(context *gin.Context, username string, password string) error {
	session := util.ContextSession(context)
	// Extract data from session
	sessionPass := session.Get("pass")
	if sessionPass != nil {
		if sessionPass != password {
			return errors.New("password error with username: " + username)
		} else {
			return nil
		}
	}
	var teamAdmin model.TeamAdmin
	result := global.DB.Where(map[string]interface{}{
		"username": username,
	}).Take(&teamAdmin)
	if result.RowsAffected == 0 {
		return errors.New("no such a volunteer")
	} else if teamAdmin.Password != password {
		// TODO md5 hash needed
		return errors.New("password error with username: " + username)
	} else {
		if teamAdmin.TeamID.Valid {
			session.Set("teamid", teamAdmin.TeamID.Int64)
		} else {
			session.Set("teamid", int64(-1))
		}
		session.Set("user", teamAdmin.Username)
		session.Set("pass", teamAdmin.Password)
		if err := session.Save(); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func ExtractAdminInfo(session sessions.Session) (model.TeamAdmin, error) {
	var admin model.TeamAdmin

	username := session.Get("user")
	if username == nil {
		return admin, errors.New("not login")
	}

	global.DB.Where(map[string]interface{}{
		"username": username,
	}).Take(&admin)
	return admin, nil
}
