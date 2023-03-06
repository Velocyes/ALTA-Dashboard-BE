package data

import (
	"gorm.io/gorm"

	"alta-dashboard-be/features/logs"
)

type Log struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Status   string `gorm:"type:enum('None', 'Join Class', 'Continue Section 2', 'Continue Section 3');default:'None';not null"`
	Feedback string `gorm:"not null"`
	UserID   uint
	MenteeID uint
}

// Mock
type Mentee struct {
	gorm.Model
	Logs []Log
}

func EntityToGorm(logEntity logs.LogEntity) Log {
	log := Log{
		Title:    logEntity.Title,
		Status:   logEntity.Status,
		Feedback: logEntity.Feedback,
		UserID:   logEntity.UserID,
		MenteeID: logEntity.MenteeID,
	}
	if logEntity.Id != 0 {
		log.ID = logEntity.Id
	}
	return log
}

func GormToEntity(logGorm Log) logs.LogEntity {
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

func ListGormToEntity(logsGorm []Log) []logs.LogEntity {
	var logEntities []logs.LogEntity
	for _, v := range logsGorm {
		logEntities = append(logEntities, GormToEntity(v))
	}
	return logEntities
}
