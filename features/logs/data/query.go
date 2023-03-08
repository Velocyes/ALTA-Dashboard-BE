package data

import (
	"alta-dashboard-be/features/logs"
	_logModel "alta-dashboard-be/features/logs/models"
	"alta-dashboard-be/utils/consts"
	"errors"
	"math"
	"strings"

	"gorm.io/gorm"
)

type logQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) logs.LogDataInterface_ {
	return &logQuery{
		db: db,
	}
}

func (logQuery *logQuery) Insert(input logs.LogEntity) (logs.LogEntity, error) {
	logGorm := EntityToGorm(input)
	txInsert := logQuery.db.Create(&logGorm)
	if txInsert.Error != nil || txInsert.RowsAffected == 0{
		if strings.Contains(txInsert.Error.Error(), "Error 1452 (23000)") {
			return logs.LogEntity{}, errors.New(consts.LOG_MenteeNotExisted)
		}
		return logs.LogEntity{}, errors.New(consts.SERVER_InternalServerError)
	}

	logEntity := GormToEntity(logGorm)
	return logEntity, nil
}

func (logQuery *logQuery) SelectData(searchedMenteeId uint, limit, offset int) (map[string]any, error) {
	logsGorm, dataCount, dataResponse := []_logModel.Log{}, int64(0), map[string]any{}
	txCount := logQuery.db.Table("logs").Where("mentee_id = ?", searchedMenteeId).Count(&dataCount)
	txSelect := logQuery.db.Model(&logsGorm).Where("mentee_id = ?", searchedMenteeId).Limit(limit).Offset(offset).Find(&logsGorm)
	if txSelect.Error != nil || txCount.Error != nil {
		return nil, errors.New(consts.SERVER_InternalServerError)
	}

	logEntities := ListGormToEntity(logsGorm)
	dataResponse["total_page"] = math.Round(float64(dataCount)/float64(int64(limit)))
	dataResponse["page"] = (offset / limit) + 1
	dataResponse["data"] = logEntities
	return dataResponse, nil
}
