package data

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Title    string `gorm:"not null;type:varchar(50)"`
	Status   string `gorm:"type:enum('None', 'Join Class', 'Continue Section 2', 'Continue Section 3');default:'None';not null"`
	Feedback string `gorm:"not null;type:text"`
	UserID   uint   `gorm:"not null;foreignKey"`
	MenteeID uint   `gorm:"not null;foreignKey"`
}
