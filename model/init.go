package model

import (
	"github.com/dump-time/games-admin-server/global"
	"github.com/dump-time/games-admin-server/log"
)

func init() {
	if err := global.DB.AutoMigrate(
		&Admin{},
		&Team{},
		&TeamAdmin{},
		&Volunteer{},
		&Job{},
	); err != nil {
		log.Error(err)
	}
}

