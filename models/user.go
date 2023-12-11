package models

import (
	"time"
)

type User struct {
	ID       []uint8 `gorm:"primaryKey"`
	TicketNo string  `json:"name"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Dept     string
	Sec      string
	Role     string `json:"role"`
	Photo    string
	//SectionID uint
	Date time.Time
	//Section   Section
	Docs []Doc
}
