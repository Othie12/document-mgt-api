package models

import "gorm.io/gorm"

type Doctype struct {
	gorm.Model
	Name      string
	Documents []Document
}
