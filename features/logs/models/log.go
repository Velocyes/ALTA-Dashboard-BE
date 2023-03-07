package data

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Status   string `gorm:"type:enum('None', 'Join Class', 'Continue Section 2', 'Continue Section 3');default:'None';not null"`
	Feedback string `gorm:"not null"`
	UserID   uint
	MenteeID uint
}

// Mock
type Mentee struct {
	gorm.Model
	Logs []Log
}
