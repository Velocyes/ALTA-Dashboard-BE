package data

import (
	"alta-dashboard-be/features/logs"
	"alta-dashboard-be/utils/consts"
	"errors"

	"gorm.io/gorm"
)

type logQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) logs.LogDataInterface {
	return &logQuery{
		db: db,
	}
}

func (logQuery *logQuery) Insert(input logs.LogEntity) (logs.LogEntity, error) {
	logGorm := EntityToGorm(input)
	txInsert := logQuery.db.Create(&logGorm)
	if txInsert.Error != nil {
		if txInsert.Error == gorm.ErrInvalidDB {
			return logs.LogEntity{}, errors.New(gorm.ErrInvalidDB.Error())
		}
		return logs.LogEntity{}, txInsert.Error
		// return logs.LogEntity{}, errors.New(consts.SERVER_InternalServerError)
	}

	logEntity := GormToEntity(logGorm)
	return logEntity, nil
}

func (logQuery *logQuery) SelectData(searchedMenteeId uint, limit, offset int) (map[string]any, error) {
	logsGorm, dataCount, dataResponse := []Log{}, int64(0), map[string]any{}
	txCount := logQuery.db.Table("logs").Where("mentee_id = ?", searchedMenteeId).Count(&dataCount)
	txSelect := logQuery.db.Model(&logsGorm).Where("mentee_id = ?", searchedMenteeId).Limit(limit).Offset(offset).Find(&logsGorm)
	if txSelect.Error != nil || txCount.Error != nil {
		if txSelect.Error == gorm.ErrInvalidDB {
			return nil, errors.New(gorm.ErrInvalidDB.Error())
		}
		return nil, errors.New(consts.SERVER_InternalServerError)
	}

	logEntities := ListGormToEntity(logsGorm)
	if limit != -1 {
		dataResponse["total_page"] = int(dataCount) / limit
		dataResponse["page"] = (offset / limit) + 1
		dataResponse["data"] = logEntities
	} else {
		dataResponse["data"] = logEntities
	}
	return dataResponse, nil
}

// func (userQuery *userQuery) UpdateData(input users.UserEntity) (users.UserEntity, error) {
// 	inputedUserGorm, updatedUserGorm := EntityToGorm(input), User{}
// 	txUpdate := userQuery.db.Model(&inputedUserGorm).Updates(inputedUserGorm)
// 	if txUpdate.Error != nil {
// 		if txUpdate.RowsAffected == 0 {
// 			return users.UserEntity{}, errors.New(consts.USER_FailedUpdate)
// 		}
// 		return users.UserEntity{}, errors.New(consts.SERVER_InternalServerError)
// 	}

// 	userQuery.db.Model(inputedUserGorm).Find(&updatedUserGorm)
// 	return GormToEntity(updatedUserGorm), nil
// }

// func (userQuery *userQuery) Delete(userId uint) error {
// 	selectedUserGorm, err := userQuery.SelectData(userId)
// 	if err != nil {
// 		return err
// 	}

// 	txDelete := userQuery.db.Model(&selectedUserGorm).Where("id = ?", userId).Delete(&selectedUserGorm)
// 	if txDelete.Error != nil {
// 		if txDelete.RowsAffected == 0 {
// 			return errors.New(consts.USER_FailedDelete)
// 		}
// 		return errors.New(consts.SERVER_InternalServerError)
// 	}
// 	return nil
// }
