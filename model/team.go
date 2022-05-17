package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model

	Name         string
	Organization string
	Code         string

	TeamAdmins []TeamAdmin `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Volunteers []Volunteer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Jobs       []Job       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
