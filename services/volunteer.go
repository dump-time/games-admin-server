package services

import (
	"database/sql"

	"github.com/dump-time/games-admin-server/global"
	"github.com/dump-time/games-admin-server/model"
)

func AddVolunteer(volunteer *model.Volunteer) error {
	result := global.DB.Create(&volunteer)
	return result.Error
}

func ListVolunteers(teamID sql.NullInt64, offset int, pageSize int) ([]model.Volunteer, error) {
	var volunteers []model.Volunteer
	result := global.DB.Debug().Where(map[string]interface{}{"team_id": teamID}).Limit(pageSize).Offset(offset).Find(&volunteers)
	return volunteers, result.Error
}
