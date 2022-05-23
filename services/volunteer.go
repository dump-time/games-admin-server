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

func SearchVolunteer(teamID sql.NullInt64, IDNumber string) (model.Volunteer, error) {
	var volunteer model.Volunteer
	result := global.DB.Where(map[string]interface{}{
		"team_id":   teamID,
		"id_number": IDNumber,
	}).Take(&volunteer)
	if result.Error != nil {
		return volunteer, result.Error
	} else if result.RowsAffected == 0 {
		return volunteer, errors.New("No such a volunteer")
	} else {
		return volunteer, nil
	}
}

func GetVolunteerPagesNum(teamID sql.NullInt64, pageSize int) (int64, error) {
	var volunteerNum int64
	global.DB.Model(&model.Volunteer{}).Where(map[string]interface{}{
		"team_id": teamID,
	}).Count(&volunteerNum)

	return (volunteerNum-1)/int64(pageSize) + 1, nil
}
