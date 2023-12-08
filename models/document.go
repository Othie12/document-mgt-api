package models

import (
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	UserID    uint64
	DoctypeID uint64
	Filepath  string
	Doctype   Doctype
	User      User
}
