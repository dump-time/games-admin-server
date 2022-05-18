package services

import (
	"github.com/dump-time/games-admin-server/global"
	"github.com/dump-time/games-admin-server/model"
)

func AddVolunteer(volunteer *model.Volunteer) error {
	result := global.DB.Create(&volunteer)
	return result.Error
}
