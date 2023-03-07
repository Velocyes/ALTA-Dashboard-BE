package data

import (
	"errors"
	"alta-dashboard-be/features/users"
	_userModel "alta-dashboard-be/features/users/models"
	"alta-dashboard-be/middlewares"
	"alta-dashboard-be/utils/consts"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

func CompareHashPassword(inputPassword, dbPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(inputPassword))
	return err == nil
}

func (userQuery *userQuery) Login(email string, password string) (users.UserEntity, string, error) {
	loggedInUserGorm := _userModel.User{}
	txSelect := userQuery.db.Where("email = ?", email).First(&loggedInUserGorm)
	if txSelect.Error != nil {
		if txSelect.Error == gorm.ErrInvalidDB {
			return users.UserEntity{}, "", errors.New(gorm.ErrInvalidDB.Error())
		}
		if txSelect.Error == gorm.ErrRecordNotFound {
			return users.UserEntity{}, "", errors.New(gorm.ErrRecordNotFound.Error())
		}
		return users.UserEntity{}, "", errors.New(consts.SERVER_InternalServerError)
	}

	if !CompareHashPassword(password, loggedInUserGorm.Password) {
		return users.UserEntity{}, "", errors.New(consts.USER_WrongPassword)
	}

	token, errToken := middlewares.CreateToken(int(loggedInUserGorm.ID), loggedInUserGorm.Role)
	if errToken != nil {
		return users.UserEntity{}, "", errToken
	}

	loggedInUserEntity := GormToEntity(loggedInUserGorm)
	return loggedInUserEntity, token, nil
}

func (userQuery *userQuery) Insert(input users.UserEntity) (users.UserEntity, error) {
	userGorm := EntityToGorm(input)
	txInsert := userQuery.db.Create(&userGorm)
	if txInsert.Error != nil {
		if txInsert.Error == gorm.ErrInvalidDB {
			return users.UserEntity{}, errors.New(gorm.ErrInvalidDB.Error())
		}
		if strings.Contains(txInsert.Error.Error(), "Error 1062 (23000)") {
			return users.UserEntity{}, errors.New(consts.USER_EmailAlreadyUsed)
		}
		return users.UserEntity{}, errors.New(consts.SERVER_InternalServerError)
	}

	userEntity := GormToEntity(userGorm)
	return userEntity, nil
}

func (userQuery *userQuery) SelectAll(limit, offset int) (map[string]any, error) {
	usersGorm, dataCount, dataResponse := []_userModel.User{}, int64(0), map[string]any{}
	txCount := userQuery.db.Table("users").Count(&dataCount)
	txSelect := userQuery.db.Limit(limit).Offset(offset).Find(&usersGorm)
	if txSelect.Error != nil || txCount.Error != nil {
		if txSelect.Error == gorm.ErrInvalidDB {
			return nil, errors.New(gorm.ErrInvalidDB.Error())
		}
		return nil, errors.New(consts.SERVER_InternalServerError)
	}

	userEntities := ListGormToEntity(usersGorm)
	if limit != -1 {
		dataResponse["total_page"] = int(dataCount) / limit
		dataResponse["page"] = (offset / limit) + 1
		dataResponse["data"] = userEntities
	} else {
		dataResponse["data"] = userEntities
	}
	return dataResponse, nil
}

func (userQuery *userQuery) SelectData(userId uint) (users.UserEntity, error) {
	userGorm := _userModel.User{}
	txSelect := userQuery.db.Find(&userGorm, userId)
	if txSelect.Error != nil {
		if txSelect.Error == gorm.ErrInvalidDB {
			return users.UserEntity{}, errors.New(gorm.ErrInvalidDB.Error())
		}
		return users.UserEntity{}, errors.New(consts.SERVER_InternalServerError)
	}

	userEntity := GormToEntity(userGorm)
	return userEntity, nil
}

func (userQuery *userQuery) UpdateData(input users.UserEntity) (users.UserEntity, error) {
	inputedUserGorm, updatedUserGorm := EntityToGorm(input), _userModel.User{}
	txUpdate := userQuery.db.Model(&inputedUserGorm).Updates(inputedUserGorm)
	if txUpdate.Error != nil {
		if txUpdate.RowsAffected == 0 {
			return users.UserEntity{}, errors.New(consts.USER_FailedUpdate)
		}
		return users.UserEntity{}, errors.New(consts.SERVER_InternalServerError)
	}

	userQuery.db.Model(inputedUserGorm).Find(&updatedUserGorm)
	return GormToEntity(updatedUserGorm), nil
}

func (userQuery *userQuery) Delete(userId uint) error {
	selectedUserGorm, err := userQuery.SelectData(userId)
	if err != nil {
		return err
	}

	txDelete := userQuery.db.Model(&selectedUserGorm).Where("id = ?", userId).Delete(&selectedUserGorm)
	if txDelete.Error != nil {
		if txDelete.RowsAffected == 0 {
			return errors.New(consts.USER_FailedDelete)
		}
		return errors.New(consts.SERVER_InternalServerError)
	}
	return nil
}
