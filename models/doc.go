package models

import (
	"gorm.io/gorm"
)

type Doc struct {
	gorm.Model
	UserID   string `json:"user_id" gorm:"type:varchar(20)"`
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Filepath string `json:"filepath" gorm:"type:varchar(255)"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL"`
}
