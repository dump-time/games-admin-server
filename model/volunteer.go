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
	Experience string `gorm:"type:text"` // 工作经历
	Status     int    // 志愿者录取状态
	Tel        string // 电话联系方式

	IntentionID sql.NullInt64 // 志愿 JobID
	Intention   Job           `gorm:"foreignKey:IntentionID"`
	JobID       sql.NullInt64
	Job         Job
}
