package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	TicketNo    string    `json:"ticket_no"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Dept        string    `json:"dept"`
	Sec         string    `json:"sec"`
	Role        string    `json:"role"`
	Designation string    `json:"designation"`
	Doj         time.Time `json:"doj"`
	Photo       string    `json:"photo"`
	SectionID   uint      `json:"section_id"`
	Date        time.Time `json:"date"`
	Section     Section   `json:"section"`
	Docs        []Doc     `json:"docs"`
}
