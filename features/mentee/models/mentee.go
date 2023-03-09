package models

import (
	cm "alta-dashboard-be/features/class/models"
	lm "alta-dashboard-be/features/logs/models"
	"database/sql"

	"gorm.io/gorm"
)

type Mentee struct {
	gorm.Model
	FullName  string    `gorm:"type:varchar(50);not null"`
	Email     string    `gorm:"type:varchar(50);not null;unique"`
	Address   string    `gorm:"type:text;not null"`
	Phone     string    `gorm:"type:varchar(15);not null"`
	Telegram  string    `gorm:"type:varchar(25);not null"`
	Emergency Emergency `gorm:"foreignKey:MenteeID;references:ID"`
	Education Education `gorm:"foreignKey:MenteeID;references:ID"`
	ClassID   int       `gorm:"not null;foreignKey"`
	Class     cm.Class  `gorm:"foreignKey:ClassID;references:ID"`
	Logs      []lm.Log
	Status    string `gorm:"type:enum('Interview', 'Join Class', 'Continue Unit 2', 'Continue Unit 3', 'Eliminated', 'Placement', 'Repeat Unit 1', 'Repeat Unit 2', 'Repeat Unit 3');default:'None';not null"`
}

type Emergency struct {
	MenteeID int    `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(50);not null"`
	Phone    string `gorm:"type:varchar(15);not null"`
	Status   string `gorm:"type:varchar(50);not null"`
}

type Education struct {
	MenteeID       int          `gorm:"primaryKey"`
	Type           string       `gorm:"type:ENUM('IT','NON-IT');not null"`
	Major          string       `gorm:"type:varchar(50);not null"`
	GraduationDate sql.NullTime `gorm:"type:date"`
}
