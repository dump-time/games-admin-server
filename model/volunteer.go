package model

import "gorm.io/gorm"

type Volunteer struct {
	gorm.Model

	TeamID     uint
	Name       string
	IDNumber   string // 身份证号
	Gender     bool
	Employment string // 目前在职情况
	Avatar     string // 头像地址
	Intention  string // 志愿
	Experience string `gorm:"type:text"` // 工作经历
	Tel        string // 电话联系方式
	JobID      uint
}
