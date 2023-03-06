package service

import (
	"errors"
	"alta-dashboard-be/features/users"
	"alta-dashboard-be/utils/consts"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userData users.UserDataInterface
	validate *validator.Validate
}

func New(userData users.UserDataInterface) users.UserServiceInterface {
	return &userService{
		userData: userData,
		validate: validator.New(),
	}
}

func HashPassword(inputPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(inputPassword), 14)
	return string(hashedPassword), err
}

func (userService *userService) Login(email string, password string) (users.UserEntity, string, error) {
	if email == "" || password == "" {
		return users.UserEntity{}, "", errors.New(consts.USER_EmptyCredentialError)
	}

	userEntity, token, err := userService.userData.Login(strings.ToLower(email), password)
	if err != nil {
		return users.UserEntity{}, "", err
	}

	return userEntity, token, err
}

func (userService *userService) Create(userInput users.UserEntity, loggedInUserRole string) (users.UserEntity, error) {
	if loggedInUserRole != consts.E_USER_Admin {
		return users.UserEntity{}, errors.New(consts.SERVER_ForbiddenRequest)
	}
	
	if userInput.Email == "" || userInput.Password == "" {
		return users.UserEntity{}, errors.New(consts.USER_EmptyCredentialError)
	}

	userInput.Email = strings.ToLower(userInput.Email)
	userInput.Password, _ = HashPassword(userInput.Password)
	err := userService.validate.Struct(userInput)
	if err != nil {
		return users.UserEntity{}, err
	}

	userEntity, err := userService.userData.Insert(userInput)
	if err != nil {
		return users.UserEntity{}, err
	}

	return userEntity, nil
}

func (userService *userService) GetAll(limit, offset int) (map[string]any, error) {
	dataResponse, err := userService.userData.SelectAll(limit, offset)
	if err != nil {
		return map[string]any{}, err
	}

	return dataResponse, err
}

func (userService *userService) GetData(loggedInUserId, userId uint, loggedInUserRole string) (users.UserEntity, error) {
	// Validasi loggedInUser melihat data siapa
	if loggedInUserId != userId && loggedInUserRole == consts.E_USER_User {
		return users.UserEntity{}, errors.New(consts.SERVER_ForbiddenRequest)
	}

	userEntity, err := userService.userData.SelectData(uint(userId))
	if err != nil {
		return users.UserEntity{}, err
	}

	return userEntity, nil
}

func (userService *userService) ModifyData(loggedInUserId, userId uint, loggedInUserRole string, userInput users.UserEntity) (users.UserEntity, error) {
	// Validasi loggedInUser mengubah data siapa
	if loggedInUserId != userId && loggedInUserRole == consts.E_USER_User {
		return users.UserEntity{}, errors.New(consts.SERVER_ForbiddenRequest)
	}

	userInput.Id = userId
	userEntity, err := userService.userData.UpdateData(userInput)
	if err != nil {
		return users.UserEntity{}, err
	}

	return userEntity, nil
}

func (userService *userService) Remove(loggedInUserId, userId uint, loggedInUserRole string) error {
	// Validasi loggedInUser menghapus data siapa
	if loggedInUserId != userId && loggedInUserRole == consts.E_USER_User {
		return errors.New(consts.SERVER_ForbiddenRequest)
	}

	err := userService.userData.Delete(userId)
	if err != nil {
		return err
	}

	return nil
}
