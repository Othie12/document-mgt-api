package models

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:@tcp(127.0.0.1:3306)/hrms?charset=utf8mb4&parseTime=True&loc=Local"
	sqldb, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("Failed to open sql db: " + err.Error())
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqldb,
	}), &gorm.Config{})

	if err != nil {
		panic("Database connection error: " + err.Error())
	}

	DB = db
}
