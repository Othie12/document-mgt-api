package models

import (
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	DeptID      uint   `json:"dept_id"`
	Document1   string `json:"document1"`
	Application string `json:"application"`
	Certificate string `json:"certificate"`
	Appointment string `json:"appointment"`
	Discipline  string `json:"discipline"`
	Others      string `json:"others"`
	//Department  Department `gorm:"foreignKey:DeptID"`
	User User `gorm:"foreignKey:UserID" json:"user"`
}
