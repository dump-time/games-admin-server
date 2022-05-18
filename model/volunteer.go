package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type Volunteer struct {
	gorm.Model

	TeamID     sql.NullInt64 `gorm:"index"`
	Name       string
	IDNumber   string // 身份证号
	Gender     bool
	Employment string // 目前在职情况
	Avatar     string // 头像地址
	Intention  int    // 志愿 JobID
	Experience string `gorm:"type:text"` // 工作经历
	Tel        string // 电话联系方式
	JobID      sql.NullInt64
}
