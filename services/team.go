package services

import (
	"errors"

	"github.com/dump-time/games-admin-server/global"
	"github.com/dump-time/games-admin-server/model"
)

func CreateTeam(team *model.Team) error {
	result := global.DB.Create(team)
	return result.Error
}

func ListTeams(offset int, pageSize int) ([]model.Team, int64, error) {
	var teams []model.Team
	var num int64

	// global.DB.Limit(pageSize).Offset(offset).Find(&teams)
	global.DB.Raw("SELECT * FROM `volunteers` WHERE `volunteers`.`deleted_at` IS NULL " +
	 "AND id >= (SELECT id FROM volunteers limit 1 OFFSET ?) LIMIT ?", offset, pageSize).Scan(&teams)
	global.DB.Model(&model.Team{}).Count(&num)

	return teams, num, nil
}

func DeleteTeam(teamID uint) error {
	result := global.DB.Where(map[string]interface{}{
		"id": teamID,
	}).Delete(&model.Team{})

	return result.Error
}

func UpdateTeam(team *model.Team) error {
	result := global.DB.Model(team).Updates(*team)
	if result.RowsAffected == 0 {
		return errors.New("no such a team")
	} else {
		return result.Error
	}
}

func GetTeamInfo(teamID uint) (model.Team, error) {
	var team model.Team
	result := global.DB.Where(map[string]interface{}{
		"id": teamID,
	}).Take(&team)

	return team, result.Error
}
