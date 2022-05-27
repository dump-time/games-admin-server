package services

import (
	"github.com/dump-time/games-admin-server/global"
	"github.com/dump-time/games-admin-server/model"
)

func CreateTeam(team *model.Team) error {
	result := global.DB.Create(team)
	return result.Error
}
