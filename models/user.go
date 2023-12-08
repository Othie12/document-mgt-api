package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Username  string
	Password  string
	SectionID string
	Role      string
	Photo     string
	Section   Section
	Documents []Document
}
