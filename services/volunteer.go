package services

import (
	"database/sql"
	"errors"

	"github.com/dump-time/games-admin-server/global"
	"github.com/dump-time/games-admin-server/model"
	"gorm.io/gorm"
)

func AddVolunteer(volunteer *model.Volunteer) error {
	result := global.DB.Create(&volunteer)
	return result.Error
}

func ListVolunteers(teamID sql.NullInt64, offset int, pageSize int) ([]model.Volunteer, error) {
	var volunteers []model.Volunteer
	var result *gorm.DB
	if teamID.Valid {
		result = global.DB.Debug().Where(
			"team_id = ? and id >= (?)", 
			teamID, 
			global.DB.Raw("SELECT id FROM volunteers WHERE team_id = ? and deleted_at IS NULL limit 1 OFFSET ?", 
				teamID.Int64, offset)).
			Preload("Intention").
			Preload("Job").
			Limit(pageSize).
			Find(&volunteers)
	} else {
		result = global.DB.Debug().Where(
			"team_id IS NULL and id >= (?)", 
			global.DB.Raw("SELECT id FROM volunteers WHERE team_id is null and deleted_at IS NULL limit 1 OFFSET ?", 
				offset)).
			Preload("Intention").
			Preload("Job").
			Limit(pageSize).
			Find(&volunteers)
	}
	
	return volunteers, result.Error
}

func DeleteVolunteer(teamID sql.NullInt64, volunteerID uint) error {
	condition := map[string]interface{}{
		"id": volunteerID,
	}
	if teamID.Valid {
		condition["team_id"] = teamID
	}
	result := global.DB.Where(condition).Delete(&model.Volunteer{})
	if result.RowsAffected == 0 {
		return errors.New("no such a volunteer in this team")
	}
	return result.Error
}

func UpdateVolunteer(teamID sql.NullInt64, volunteerID uint, volunteer *model.Volunteer) error {
	condition := map[string]interface{}{
		"id": volunteerID,
	}
	if teamID.Valid {
		condition["team_id"] = teamID
	}
	result := global.DB.Model(&model.Volunteer{}).Where(condition).Updates(map[string]interface{}{
		"name":       volunteer.Name,
		"gender":     volunteer.Gender,
		"job_id":     volunteer.JobID,
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
		return errors.New("no such a volunteer in this team")
	}
	return nil
}

// SearchVolunteer get volunteer data with user id card number
func SearchVolunteer(teamID sql.NullInt64, IDNumber string) (model.Volunteer, error) {
	var volunteer model.Volunteer
	condition := map[string]interface{}{
		"id_number": IDNumber,
		"team_id":   teamID,
	}
	result := global.DB.Where(condition).Preload("Intention").Preload("Job").Take(&volunteer)
	if result.Error != nil {
		return volunteer, result.Error
	} else if result.RowsAffected == 0 {
		return volunteer, errors.New("no such a volunteer")
	} else {
		return volunteer, nil
	}
}

// GetVolunteersNum fetch all tuple's count in the table
func GetVolunteersNum(teamID sql.NullInt64, pageSize int) (int64, error) {
	var volunteersNum int64
	condition := map[string]interface{}{}
	if teamID.Valid {
		condition["team_id"] = teamID.Int64
	}
	global.DB.Model(&model.Volunteer{}).Where(condition).Count(&volunteersNum)

	return volunteersNum, nil
}
