package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type Job struct {
	gorm.Model

	TeamID sql.NullInt64
	Name string
	Content string
	Location string
	Volunteers []Volunteer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
