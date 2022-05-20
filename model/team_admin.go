package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type TeamAdmin struct {
	gorm.Model

	Username string
	Password string
	TeamID   sql.NullInt64
}
