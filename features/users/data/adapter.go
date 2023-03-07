package data

import (
	"alta-dashboard-be/features/users"
	_userModel "alta-dashboard-be/features/users/models"
)

func EntityToGorm(userEntity users.UserEntity) _userModel.User {
	userGorm := _userModel.User{
		FullName: userEntity.FullName,
		Email:    userEntity.Email,
		Password: userEntity.Password,
		Team:     userEntity.Team,
		Role:     userEntity.Role,
		Status:   userEntity.Status,
	}
	if userEntity.Id != 0 {
		userGorm.ID = userEntity.Id
	}
	return userGorm
}

func GormToEntity(userGorm _userModel.User) users.UserEntity {
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

func ListGormToEntity(usersGorm []_userModel.User) []users.UserEntity {
	var userEntities []users.UserEntity
	for _, v := range usersGorm {
		userEntities = append(userEntities, GormToEntity(v))
	}
	return userEntities
}
