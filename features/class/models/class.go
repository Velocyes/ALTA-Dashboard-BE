package models

import (
	u "alta-dashboard-be/features/users/data"
	"time"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name      string    `gorm:"not null;type:varchar(50)"`
	ShortName string    `gorm:"not null;type:varchar(10)"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
	UserID    int       `gorm:"not null";foreignKey`
	User      u.User    `foreignKey:UserID,references:ID`
}
