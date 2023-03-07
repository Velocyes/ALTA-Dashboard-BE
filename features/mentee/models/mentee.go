package models

import (
	cm "alta-dashboard-be/features/class/models"
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
	Emergency Emergency `gorm:"foreignKey:ID;references:MenteeID"`
	Education Education `gorm:"foreignKey:ID;references:MenteeID"`
	ClassID   int
	Class     cm.Class `gorm:"foreignKey:ClassID;references:ID"`
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
