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

	global.DB.Debug().Limit(pageSize).Offset(offset).Find(&teams)
	global.DB.Model(&model.Team{}).Count(&num)

	return teams, num, nil
}
