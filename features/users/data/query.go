package data

import (
	"alta-dashboard-be/features/users"
	_userModel "alta-dashboard-be/features/users/models"
	"alta-dashboard-be/middlewares"
	"alta-dashboard-be/utils/consts"
	"errors"
	"fmt"
	"math"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.UserDataInterface_ {
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
		if strings.Contains(txInsert.Error.Error(), "Error 1062 (23000)") {
			return users.UserEntity{}, errors.New(consts.USER_EmailAlreadyUsed)
		}
		return users.UserEntity{}, errors.New(consts.SERVER_InternalServerError)
	}

	userEntity := GormToEntity(userGorm)
	return userEntity, nil
}

func (userQuery *userQuery) SelectAll(queryParams map[string]any, limit, offset int) (map[string]any, error) {
	usersGorm, dataCount, dataResponse := []_userModel.User{}, int64(0), map[string]any{}
	txCount := userQuery.db.Table("users").Where(queryParams).Count(&dataCount)
	txSelect := userQuery.db.Where(queryParams).Limit(limit).Offset(offset).Find(&usersGorm)
	if txSelect.Error != nil || txCount.Error != nil {
		if strings.Contains(txSelect.Error.Error(), "Error 1054 (42S22)") {
			return nil, errors.New(consts.DATABASE_InvalidQueryParameter)
		}
		return nil, errors.New(consts.SERVER_InternalServerError)
	}

	userEntities := ListGormToEntity(usersGorm)
	dataResponse["total_page"] = math.Round(float64(dataCount)/float64(int64(limit)))
	dataResponse["page"] = (offset / limit) + 1
	dataResponse["data"] = userEntities
	return dataResponse, nil
}

func (userQuery *userQuery) SelectData(userId uint) (users.UserEntity, error) {
	userGorm := _userModel.User{}
	txSelect := userQuery.db.Model(&userGorm).Where("id = ?", userId).First(&userGorm)
	if txSelect.Error != nil {
		if txSelect.Error == gorm.ErrRecordNotFound {
			return users.UserEntity{}, errors.New(gorm.ErrRecordNotFound.Error())
		}
		return users.UserEntity{}, errors.New(consts.SERVER_InternalServerError)
	}

	userEntity := GormToEntity(userGorm)
	return userEntity, nil
}

func (userQuery *userQuery) UpdateData(input users.UserEntity) (users.UserEntity, error) {
	inputedUserGorm, updatedUserGorm := EntityToGorm(input), _userModel.User{}
	txUpdate := userQuery.db.Model(&inputedUserGorm).Updates(&inputedUserGorm)
	if txUpdate.Error != nil {
		if strings.Contains(txUpdate.Error.Error(), "Error 1062 (23000)") {
			return users.UserEntity{}, errors.New(consts.USER_EmailAlreadyUsed)
		}
		return users.UserEntity{}, errors.New(consts.SERVER_InternalServerError)
	}
	if txUpdate.RowsAffected == 0 {
		return users.UserEntity{}, errors.New(gorm.ErrRecordNotFound.Error())
	}

	fmt.Println(txUpdate.RowsAffected)

	userQuery.db.Find(&updatedUserGorm, "id = ?", inputedUserGorm.ID)
	return GormToEntity(updatedUserGorm), nil
}

func (userQuery *userQuery) Delete(userId uint) error {
	txDelete := userQuery.db.Model(&_userModel.User{}).Where("id = ?", userId).Delete(&_userModel.User{})
	if txDelete.Error != nil{
		return errors.New(consts.SERVER_InternalServerError)
	}
	if txDelete.RowsAffected == 0 {
		return errors.New(gorm.ErrRecordNotFound.Error())
	}
	return nil
}
