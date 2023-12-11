package models

import "gorm.io/gorm"

type Doc struct {
	gorm.Model
	UserID   uint
	Name     string
	Filepath string
	User     User `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL"`
}
