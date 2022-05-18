package services

import (
	"database/sql"
	"github.com/dump-time/games-admin-server/global"
	"github.com/dump-time/games-admin-server/model"
)

func AddJob(job *model.Job) error {
	return global.DB.Create(job).Error
}

func GetJobs(teamId sql.NullInt64) ([]model.Job, error) {
	var results []model.Job
	err := global.DB.Find(&results, model.Job{
		TeamID: teamId,
	}).Error

	return results, err
}
