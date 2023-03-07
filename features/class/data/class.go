package data

import (
	"alta-dashboard-be/features/class"
	"alta-dashboard-be/features/class/models"
	"alta-dashboard-be/utils/helper"
	"errors"

	"gorm.io/gorm"
)

type ClassData struct {
	db *gorm.DB
}

// Create implements class.ClassData_
func (u *ClassData) Create(class class.ClassCore) error {
	mdl := convertToModel(&class)
	tx := u.db.Create(&mdl)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return errors.New("error in database")
	}
	return nil
}

// Delete implements class.ClassData_
func (u *ClassData) Delete(id int) error {
	var mdl models.Class
	tx := u.db.Where("id = ?", id).Delete(&mdl)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return errors.New("error in database")
	}
	return nil
}

// GetAll implements class.ClassData_
func (u *ClassData) GetAll(page int, limit int) ([]class.ClassCore, error) {
	mdl := []models.Class{}
	limit, offset := helper.LimitOffsetConvert(page, limit)
	tx := u.db.Offset(offset).Limit(limit).Find(&mdl)
	if tx.Error != nil {
		return nil, errors.New("error in database")
	}
	return convertToCoreList(mdl), nil
}

// GetOne implements class.ClassData_
func (u *ClassData) GetOne(id int) (class.ClassCore, error) {
	mdl := models.Class{}
	tx := u.db.Where("id = ?", id).First(&mdl)
	if tx.Error != nil {
		return class.ClassCore{}, errors.New("error in database")
	}
	return convertToCore(&mdl), nil
}

// Update implements class.ClassData_
func (u *ClassData) Update(id int, class class.ClassCore) error {
	mdl := convertToModel(&class)
	tx := u.db.Where("id = ?", id).Updates(&mdl)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return errors.New("error in database")
	}
	return nil
}

func New(db *gorm.DB) class.ClassData_ {
	return &ClassData{
		db: db,
	}
}
