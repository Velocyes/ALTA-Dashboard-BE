package service

import (
	"alta-dashboard-be/features/class"
	"alta-dashboard-be/utils/consts"
	"errors"
)

type ClassService struct {
	data class.ClassData_
}

// Create implements class.ClassService_
func (u *ClassService) Create(userID int, class class.ClassCore) error {
	//validate userID
	if userID <= 0 {
		return errors.New(consts.VALIDATION_userID)
	}
	//set userID as userID from jwt
	class.UserID = userID
	//validate
	err := validate(&class)
	if err != nil {
		return err
	}
	return u.data.Create(class)
}

// Delete implements class.ClassService_
func (u *ClassService) Delete(userID int, id int) error {
	//validate UserID
	if userID <= 0 || id <= 0 {
		return errors.New(consts.VALIDATION_id_userID)
	}
	return u.data.Delete(id)
}

// GetAll implements class.ClassService_
func (u *ClassService) GetAll(page int, limit int) ([]class.ClassCore, error) {
	//validate page and limit
	if page <= 0 || limit <= 0 {
		return nil, errors.New(consts.VALIDATION_page_limit)
	}
	return u.data.GetAll(page, limit)
}

// GetOne implements class.ClassService_
func (u *ClassService) GetOne(id int) (class.ClassCore, error) {
	//validate id
	if id <= 0 {
		return class.ClassCore{}, errors.New(consts.VALIDATION_id)
	}
	return u.data.GetOne(id)
}

// Update implements class.ClassService_
func (u *ClassService) Update(userID int, id int, class class.ClassCore) error {
	//validate userID
	if userID <= 0 || id <= 0 {
		return errors.New(consts.VALIDATION_id_userID)
	}
	//set userID from JWT
	class.UserID = userID
	//validate
	err := validate(&class)
	if err != nil {
		return err
	}
	return u.data.Update(id, class)
}

func New(data class.ClassData_) class.ClassService_ {
	return &ClassService{
		data: data,
	}
}
