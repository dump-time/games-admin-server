package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model

	Username string
	Password string `gorm:"type:varchar(32)"` // md5 encryped	
}