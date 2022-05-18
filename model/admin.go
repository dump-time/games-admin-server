package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model

	Username string
	Password string `gorm:"type:char(32)"` // md5 encrypted
}
