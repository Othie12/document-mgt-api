package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:@tcp(127.0.0.1:3306)/hrms_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connection error: " + err.Error())
	}

	err = db.AutoMigrate(&Department{}, &Doctype{}, &Section{}, &User{}, &Document{})
	if err != nil {
		panic("Migration failed: " + err.Error())
	}

	DB = db
}
