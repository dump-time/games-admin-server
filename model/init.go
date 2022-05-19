package model

import (
	"github.com/dump-time/games-admin-server/global"
	"github.com/dump-time/games-admin-server/log"
)

func init() {
	err := global.DB.AutoMigrate(
		&Team{},
		&TeamAdmin{},
		&Volunteer{},
		&Job{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
