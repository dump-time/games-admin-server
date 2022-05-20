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

func UpdateVolunteer(teamID sql.NullInt64, volunteerID uint, volunteer *model.Volunteer) error {
	result := global.DB.Model(&model.Volunteer{}).Where(map[string]interface{}{
		"id":      volunteerID,
		"team_id": teamID,
	}).Updates(map[string]interface{}{
		"name":       volunteer.Name,
		"gender":     volunteer.Gender,
		"intention":  volunteer.Intention,
		"tel":        volunteer.Tel,
		"experience": volunteer.Experience,
		"avatar":     volunteer.Avatar,
		"id_number":  volunteer.IDNumber,
		"employment": volunteer.Employment,
		"team_id":    volunteer.TeamID,
		"status":     volunteer.Status,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("No such a volunteer in this team")
	}
	return nil
}
