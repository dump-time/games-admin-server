package services

import (
	"github.com/dump-time/games-admin-server/global"
	"github.com/dump-time/games-admin-server/model"
)

func AddJob(job *model.Job) error {
	return global.DB.Create(job).Error
}
