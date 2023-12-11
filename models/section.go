package models

import "gorm.io/gorm"

type Section struct {
	gorm.Model
	Name         string
	DepartmentID uint
	Department   Department
	//Users        []User
}
