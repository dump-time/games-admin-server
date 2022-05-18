package model

import "gorm.io/gorm"

type Job struct {
	gorm.Model

	TeamID *uint
	Name string
	Content string
	Location string
	Volunteers []Volunteer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
