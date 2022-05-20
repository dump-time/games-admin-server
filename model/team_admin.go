package model

import "gorm.io/gorm"

type TeamAdmin struct {
	gorm.Model

	username string
	password string
	TeamID uint
}
