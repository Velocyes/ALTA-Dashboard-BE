package models

import (
	_classModel "alta-dashboard-be/features/class/models"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `gorm:"not null;type:varchar(50)"`
	Email    string `gorm:"not null;unique;type:varchar(50)"`
	Password string `gorm:"not null;type:text"`
	Team     string `gorm:"not null;type:enum('Mentor', 'Placement', 'People', 'Admission', 'Academic');default:'Mentor'"`
	Role     string `gorm:"not null;type:enum('User', 'Admin');default:'User'"`
	Status   string `gorm:"not null;type:enum('Active', 'Not-Active', 'Deleted');default:'Active'"`
	Classes  []_classModel.Class
}
