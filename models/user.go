package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          string `gorm:"type:varchar(20);not null;primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	TicketNo    string         `json:"ticket_no"`
	Username    string         `json:"username"`
	Password    string         `json:"password"`
	Dept        string         `json:"dept"`
	Section     string         `json:"section" gorm:"type:varchar(20)"`
	Role        string         `json:"role"`
	Designation string         `json:"designation"`
	Doj         time.Time      `json:"doj"`
	Photo       string         `json:"photo"`
	Docs        []Doc          `json:"docs"`
}
