package services

import (
	"errors"

	"github.com/dump-time/games-admin-server/global"
	"github.com/dump-time/games-admin-server/model"
	"github.com/dump-time/games-admin-server/util"
	"github.com/gin-gonic/gin"
)

func CheckAuth(context *gin.Context, username string, password string) error {
	session := util.ContextSession(context)
	// Extract data from session
	sessionPass := session.Get(username)
	if sessionPass != nil {
		if sessionPass != password {
			return errors.New("Password error with username: " + username)
		} else {
			return nil
		}
	}

	var teamAdmin model.TeamAdmin
	result := global.DB.Where(map[string]interface{}{
		"username": username,
	}).Take(&teamAdmin)
	if result.RowsAffected == 0 {
		return errors.New("No such a volunteer")
	} else if teamAdmin.Password != password {
		// TODO md5 hash needed
		return errors.New("Password error with username: " + username)
	} else {
		session.Set(teamAdmin.Username, teamAdmin.Password)
		if err := session.Save(); err != nil {
			return err
		} else {
			return nil
		}
	}
}
