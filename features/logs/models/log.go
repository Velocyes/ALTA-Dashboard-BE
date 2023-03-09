package data

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Title    string `gorm:"not null;type:varchar(50)"`
	Status   string `gorm:"not null;type:enum('Continue Unit 1','Continue Section 2', 'Continue Section 3', 'Eliminated', 'Interview', 'Graduated', 'Placement', 'Repeat Unit 1', 'Repeat Unit 2', 'Repeat Unit 3');default:'Continue Unit 1'"`
	Feedback string `gorm:"not null;type:text"`
	UserID   uint   `gorm:"not null;foreignKey"`
	MenteeID uint   `gorm:"not null;foreignKey"`
}
