package data

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Title    string `gorm:"not null;type:varchar(50)"`
	Status   string `gorm:"type:enum('Interview', 'Join Class', 'Continue Unit 2', 'Continue Unit 3', 'Eliminated', 'Placement', 'Repeat Unit 1', 'Repeat Unit 2', 'Repeat Unit 3');default:'None';not null"`
	Feedback string `gorm:"not null;type:text"`
	UserID   uint   `gorm:"not null;foreignKey"`
	MenteeID uint   `gorm:"not null;foreignKey"`
}
