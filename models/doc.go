package models

import "gorm.io/gorm"

type Doc struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	Name     string `json:"name"`
	Filepath string `json:"filepath"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL"`
}
