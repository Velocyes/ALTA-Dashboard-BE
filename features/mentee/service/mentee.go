package service

import (
	"alta-dashboard-be/features/mentee"
	"errors"
	"net/url"
)

type MenteeService struct {
	data mentee.MenteeData_
}

// Create implements mentee.MenteeService_
func (u *MenteeService) Create(userID int, mentee mentee.MenteeCore) error {
	//set mentee status to join class
	mentee.Status = "Join Class"
	//validate userID
	if userID <= 0 {
		return errors.New("invalid userID")
	}
	//validate
	err := validate(&mentee)
	if err != nil {
		return err
	}
	return u.data.Create(mentee)
}

// Delete implements mentee.MenteeService_
func (u *MenteeService) Delete(userID int, id int) error {
	//validate userID or id
	if userID <= 0 || id <= 0 {
		return errors.New("invalid userID or id")
	}
	return u.data.Delete(id)
}

// GetAll implements mentee.MenteeService_
func (u *MenteeService) GetAll(page int, limit int, status string) ([]mentee.MenteeCore, error) {
	//validate page and limit
	if page <= 0 || limit <= 0 {
		return nil, errors.New("invalid page or limit given")
	}
	//if status is not nil, use getallfilteredbystatus
	if status != "" {
		//decode url query params
		decodedValue, err := url.QueryUnescape(status)
		if err != nil {
			return nil, errors.New("url encoded invalid")
		}
		//set status = devocedValue
		status = decodedValue
		//validate
		err = validateStatus(status)
		if err != nil {
			return nil, err
		}
		return u.data.GetAllFilteredByStatus(page, limit, status)
	}
	return u.data.GetAll(page, limit)
}

// GetOne implements mentee.MenteeService_
func (u *MenteeService) GetOne(id int) (mentee.MenteeCore, error) {
	//validate id
	if id <= 0 {
		return mentee.MenteeCore{}, errors.New("invalid id")
	}
	return u.data.GetOne(id)
}

// Update implements mentee.MenteeService_
func (u *MenteeService) Update(userID int, id int, mentee mentee.MenteeCore) error {
	//validate userID or id
	if userID <= 0 || id <= 0 {
		return errors.New("invalid userID or id")
	}
	//validate
	err := validate(&mentee)
	if err != nil {
		return err
	}
	return u.data.Update(id, mentee)
}

func New(data mentee.MenteeData_) mentee.MenteeService_ {
	return &MenteeService{
		data: data,
	}
}
