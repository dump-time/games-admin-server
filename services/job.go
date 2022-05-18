package services

import (
	"database/sql"
	"github.com/dump-time/games-admin-server/global"
	"github.com/dump-time/games-admin-server/model"
	"time"
)

func AddJob(job *model.Job) error {
	return global.DB.Create(job).Error
}

type Jobs []struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	Location  string    `json:"location"`
}

func GetJobs(teamId sql.NullInt64) (Jobs, error) {
	var results Jobs
	err := global.DB.Model(model.Job{}).Find(&results, model.Job{
		TeamID: teamId,
	}).Error

	return results, err
}
