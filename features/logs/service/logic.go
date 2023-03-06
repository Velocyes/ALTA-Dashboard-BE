package service

import (
	"alta-dashboard-be/features/logs"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type logService struct {
	logData  logs.LogDataInterface
	validate *validator.Validate
}

func New(logData logs.LogDataInterface) logs.LogServiceInterface {
	return &logService{
		logData:  logData,
		validate: validator.New(),
	}
}

func HashPassword(inputPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(inputPassword), 14)
	return string(hashedPassword), err
}

func (logService *logService) Create(logInput logs.LogEntity, loggedInUserId uint) (logs.LogEntity, error) {
	logInput.UserID = loggedInUserId
	err := logService.validate.Struct(logInput)
	if err != nil {
		return logs.LogEntity{}, err
	}

	logEntity, err := logService.logData.Insert(logInput)
	if err != nil {
		return logs.LogEntity{}, err
	}

	return logEntity, nil
}

func (logService *logService) GetData(searchedMenteeId uint, limit, offset int) (map[string]any, error) {
	dataResponse, err := logService.logData.SelectData(searchedMenteeId, limit, offset)
	if err != nil {
		return map[string]any{}, err
	}

	return dataResponse, nil
}

// func (userService *userService) ModifyData(loggedInUserId, userId uint, loggedInUserRole string, userInput users.UserEntity) (users.UserEntity, error) {
// 	// Validasi loggedInUser mengubah data siapa
// 	if loggedInUserId != userId && loggedInUserRole == consts.E_USER_User {
// 		return users.UserEntity{}, errors.New(consts.SERVER_ForbiddenRequest)
// 	}

// 	userInput.Id = userId
// 	userEntity, err := userService.userData.UpdateData(userInput)
// 	if err != nil {
// 		return users.UserEntity{}, err
// 	}

// 	return userEntity, nil
// }

// func (userService *userService) Remove(loggedInUserId, userId uint, loggedInUserRole string) error {
// 	// Validasi loggedInUser menghapus data siapa
// 	if loggedInUserId != userId && loggedInUserRole == consts.E_USER_User {
// 		return errors.New(consts.SERVER_ForbiddenRequest)
// 	}

// 	err := userService.userData.Delete(userId)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
