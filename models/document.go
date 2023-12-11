package models

import (
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	UserID      int
	DeptID      uint
	Document1   string
	Application string
	Certificate string
	Appointment string
	Discipline  string
	Others      string
	//	Department  Department `gorm:"foreignKey:DeptID"`
	//	User        User       `gorm:"foreignKey:UserID"`
}
