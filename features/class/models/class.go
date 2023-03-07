package models

import (
	"time"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name      string    `gorm:"not null;type:varchar(50)"`
	ShortName string    `gorm:"not null;type:varchar(10)"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
	UserID    int       `gorm:"not null;foreignKey"`
	// User      u.User    `gorm:"foreignKey:UserID,references:ID"`
}
