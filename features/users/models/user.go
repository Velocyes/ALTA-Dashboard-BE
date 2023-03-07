package models

import (
	"gorm.io/gorm"
	_logModel "alta-dashboard-be/features/logs/models"
)

type User struct {
	gorm.Model
	FullName string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Team     string `gorm:"type:enum('Mentor', 'Placement', 'People', 'Admission', 'Academic');default:'Mentor';not null"`
	Role     string `gorm:"type:enum('User', 'Admin');default:'User';not null"`
	Status   string `gorm:"type:enum('Active', 'Not-Active', 'Deleted');default:'Active';not null"`
	Logs     []_logModel.Log
}