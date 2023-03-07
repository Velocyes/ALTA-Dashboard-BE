package data

import (
	"alta-dashboard-be/features/logs"
	_logModel "alta-dashboard-be/features/logs/models"
)

func EntityToGorm(logEntity logs.LogEntity) _logModel.Log {
	logGorm := _logModel.Log{
		Title:    logEntity.Title,
		Status:   logEntity.Status,
		Feedback: logEntity.Feedback,
		UserID:   logEntity.UserID,
		MenteeID: logEntity.MenteeID,
	}
	if logEntity.Id != 0 {
		logGorm.ID = logEntity.Id
	}
	return logGorm
}

func GormToEntity(logGorm _logModel.Log) logs.LogEntity {
	return logs.LogEntity{
		Id:        logGorm.ID,
		Title:     logGorm.Title,
		Status:    logGorm.Status,
		Feedback:  logGorm.Feedback,
		UserID:    logGorm.UserID,
		MenteeID:  logGorm.MenteeID,
		CreatedAt: logGorm.CreatedAt,
		UpdatedAt: logGorm.UpdatedAt,
	}
}

func ListGormToEntity(logsGorm []_logModel.Log) []logs.LogEntity {
	var logEntities []logs.LogEntity
	for _, v := range logsGorm {
		logEntities = append(logEntities, GormToEntity(v))
	}
	return logEntities
}