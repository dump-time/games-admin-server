package model

import "github.com/dump-time/games-admin-server/global"

func init() {
	global.DB.AutoMigrate(
		&Admin{},
		&Team{},
		&TeamAdmin{},
		&Volunteer{},
		&Job{},
	)
}