package services

import (
	"database/sql"
	"errors"

	"github.com/dump-time/games-admin-server/global"
	"github.com/dump-time/games-admin-server/model"
)

func AddVolunteer(volunteer *model.Volunteer) error {
	result := global.DB.Create(&volunteer)
	return result.Error
}

func ListVolunteers(teamID sql.NullInt64, offset int, pageSize int) ([]model.Volunteer, error) {
	var volunteers []model.Volunteer
	result := global.DB.Where(map[string]interface{}{"team_id": teamID}).
		Limit(pageSize).Offset(offset).
		Find(&volunteers)
	return volunteers, result.Error
}

func DeleteVolunteer(teamID sql.NullInt64, volunteerID uint) error {
	result := global.DB.Where(map[string]interface{}{
		"id":      volunteerID,
		"team_id": teamID,
	}).Delete(&model.Volunteer{})
	if result.RowsAffected == 0 {
		return errors.New("No such a volunteer in this team")
	}
	return result.Error
}
