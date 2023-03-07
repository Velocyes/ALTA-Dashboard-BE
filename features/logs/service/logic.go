package service

import (
	"alta-dashboard-be/features/logs"

	"github.com/go-playground/validator/v10"
)

type logService struct {
	logData  logs.LogDataInterface_
	validate *validator.Validate
}

func New(logData logs.LogDataInterface_) logs.LogServiceInterface_ {
	return &logService{
		logData:  logData,
		validate: validator.New(),
	}
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
