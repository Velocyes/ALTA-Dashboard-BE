package data

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Title    string `gorm:"not null;type:varchar(50)"`
	Status   string `gorm:"not null;type:enum('Interview','Join Class','Unit 1','Unit 2','Unit 3','Repeat Unit 1','Repeat Unit 2','Repeat Unit 3','Placement','Eliminated','Graduated');default:'Join Class'"`
	Feedback string `gorm:"not null;type:text"`
	UserID   uint   `gorm:"not null;foreignKey"`
	MenteeID uint   `gorm:"not null;foreignKey"`
}
