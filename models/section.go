package models

import "gorm.io/gorm"

type Section struct {
	gorm.Model
	Name         string
	DepartmentID uint64
	Department   Department
	Users        []User
}
