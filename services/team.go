package services

import (
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

	global.DB.Limit(pageSize).Offset(offset).Find(&teams)
	global.DB.Model(&model.Team{}).Count(&num)

	return teams, num, nil
}

func DeleteTeam(teamID uint) error {
	result := global.DB.Where(map[string]interface{}{
		"id": teamID,
	}).Delete(&model.Team{})

	return result.Error
}
