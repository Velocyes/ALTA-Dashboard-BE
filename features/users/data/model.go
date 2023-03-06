package data

import (
	"gorm.io/gorm"

	"alta-dashboard-be/features/users"
)

type User struct {
	gorm.Model
	FullName string
	Email    string `gorm:"unique"`
	Password string
	Team     string `gorm:"type:enum('Mentor', 'Placement', 'People', 'Admission', 'Academic');default:'Mentor';not null"`
	Role     string `gorm:"type:enum('User', 'Admin');default:'User';not null"`
	Status   string `gorm:"type:enum('Active', 'Not-Active', 'Deleted');default:'Active';not null"`
}

func EntityToGorm(userEntity users.UserEntity) User {
	user := User{
		FullName: userEntity.FullName,
		Email:    userEntity.Email,
		Password: userEntity.Password,
		Team:     userEntity.Team,
		Role:     userEntity.Role,
		Status:   userEntity.Status,
	}
	if userEntity.Id != 0 {
		user.ID = userEntity.Id
	}
	return user
}

func GormToEntity(userGorm User) users.UserEntity {
	return users.UserEntity{
		Id:        userGorm.ID,
		FullName:  userGorm.FullName,
		Email:     userGorm.Email,
		Password:  userGorm.Password,
		Team:      userGorm.Team,
		Role:      userGorm.Role,
		Status:    userGorm.Status,
		CreatedAt: userGorm.CreatedAt,
		UpdatedAt: userGorm.UpdatedAt,
	}
}

func ListGormToEntity(usersGorm []User) []users.UserEntity {
	var userEntities []users.UserEntity
	for _, v := range usersGorm {
		userEntities = append(userEntities, GormToEntity(v))
	}
	return userEntities
}
