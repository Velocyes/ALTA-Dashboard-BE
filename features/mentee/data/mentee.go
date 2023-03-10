package data

import (
	"alta-dashboard-be/features/mentee"
	"alta-dashboard-be/features/mentee/models"
	"alta-dashboard-be/utils/consts"
	"alta-dashboard-be/utils/helper"
	"errors"

	"gorm.io/gorm"
)

type MenteeData struct {
	db *gorm.DB
}

// GetAllFilteredByStatus implements mentee.MenteeData_
func (u *MenteeData) GetAllFilteredByStatus(page int, limit int, status string) ([]mentee.MenteeCore, error) {
	mdl := []models.Mentee{}
	limit, offset := helper.LimitOffsetConvert(page, limit)
	tx := u.db.Where("status = ?", status).Offset(offset).Limit(limit).Preload("Education").Find(&mdl)
	if tx.Error != nil {
		return nil, errors.New(consts.DATABASE_internal_error)
	}
	return convertToCoreList(mdl), nil
}

// Create implements mentee.MenteeData_
func (u *MenteeData) Create(mentee mentee.MenteeCore) error {
	//start transaction
	tx := u.db.Begin()
	if tx.Error != nil {
		tx.Rollback()
		return errors.New(consts.DATABASE_internal_error)
	}

	//convert core to models
	m, emergency, edu := convertToModels(&mentee)

	//insert to mentee
	txC := tx.Create(&m)
	if txC.Error != nil || txC.RowsAffected == 0 {
		tx.Rollback()
		return errors.New(consts.DATABASE_internal_error)
	}

	//set menteeID to lastInsertedID
	if m.ID == 0 {
		tx.Rollback()
		return errors.New(consts.DATABASE_internal_error)
	}
	emergency.MenteeID = int(m.ID)
	edu.MenteeID = int(m.ID)

	//insert to emergency
	txC = tx.Create(&emergency)
	if txC.Error != nil || txC.RowsAffected == 0 {
		tx.Rollback()
		return errors.New(consts.DATABASE_internal_error)
	}

	//insert to education
	txC = tx.Create(&edu)
	if txC.Error != nil || txC.RowsAffected == 0 {
		tx.Rollback()
		return errors.New(consts.DATABASE_internal_error)
	}

	txC = tx.Commit()
	if txC.Error != nil {
		tx.Rollback()
		return errors.New(consts.DATABASE_commit_error)
	}

	return nil
}

// Delete implements mentee.MenteeData_
func (u *MenteeData) Delete(id int) error {
	var mdl models.Mentee
	tx := u.db.Where("id = ?", id).Delete(&mdl)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return errors.New(consts.DATABASE_internal_error)
	}
	return nil
}

// GetAll implements mentee.MenteeData_
func (u *MenteeData) GetAll(page int, limit int) ([]mentee.MenteeCore, error) {
	mdl := []models.Mentee{}
	limit, offset := helper.LimitOffsetConvert(page, limit)
	tx := u.db.Offset(offset).Limit(limit).Preload("Education").Find(&mdl)
	if tx.Error != nil {
		return nil, errors.New(consts.DATABASE_internal_error)
	}
	return convertToCoreList(mdl), nil
}

// GetOne implements mentee.MenteeData_
func (u *MenteeData) GetOne(id int) (mentee.MenteeCore, error) {
	mdl := models.Mentee{}
	tx := u.db.Where("id = ?", id).Preload("Emergency").Preload("Education").First(&mdl)
	if tx.Error != nil {
		return mentee.MenteeCore{}, errors.New(consts.DATABASE_internal_error)
	}
	return convertToCore(&mdl), nil
}

// Update implements mentee.MenteeData_
func (u *MenteeData) Update(id int, mentee mentee.MenteeCore) error {
	//start transaction
	tx := u.db.Begin()
	if tx.Error != nil {
		tx.Rollback()
		return errors.New(consts.DATABASE_internal_error)
	}

	//convert core to models
	m, emergency, edu := convertToModels(&mentee)

	//update to mentee
	txU := tx.Where("id = ?", id).Updates(&m)
	if txU.Error != nil || txU.RowsAffected == 0 {
		tx.Rollback()
		return errors.New(consts.DATABASE_internal_error)
	}

	//set menteeID to lastInsertedID
	if m.ID == 0 {
		tx.Rollback()
		return errors.New(consts.DATABASE_internal_error)
	}
	emergency.MenteeID = int(m.ID)
	edu.MenteeID = int(m.ID)

	//update to emergency
	txU = tx.Where("mentee_id = ?", emergency.MenteeID).Updates(&emergency)
	if txU.Error != nil || txU.RowsAffected == 0 {
		tx.Rollback()
		return errors.New(consts.DATABASE_internal_error)
	}

	//update to education
	txU = tx.Where("mentee_id = ?", edu.MenteeID).Updates(&edu)
	if txU.Error != nil || txU.RowsAffected == 0 {
		tx.Rollback()
		return errors.New(consts.DATABASE_internal_error)
	}

	txU = tx.Commit()
	if txU.Error != nil {
		tx.Rollback()
		return errors.New(consts.DATABASE_commit_error)
	}

	return nil
}

func New(db *gorm.DB) mentee.MenteeData_ {
	return &MenteeData{
		db: db,
	}
}
