package service

import (
	"alta-dashboard-be/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	table := CreateTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock data
			menteeData := new(mocks.MenteeData_)
			menteeData.On("Create", mock.Anything).Return(nil)

			//execute
			service := New(menteeData)
			err := service.Create(v.Input.UserID, v.Input.Mentee)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	table := UpdateTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock data
			menteeData := new(mocks.MenteeData_)
			menteeData.On("Update", mock.AnythingOfType("int"), mock.Anything).Return(nil)

			//execute
			service := New(menteeData)
			err := service.Update(v.Input.UserID, v.Input.ID, v.Input.Mentee)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	table := DeleteTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock data
			menteeData := new(mocks.MenteeData_)
			menteeData.On("Delete", mock.AnythingOfType("int")).Return(nil)

			//execute
			service := New(menteeData)
			err := service.Delete(v.Input.UserID, v.Input.ID)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	table := GetAllTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock data
			menteeData := new(mocks.MenteeData_)
			menteeData.On("GetAll", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(v.Output.Result, nil)
			menteeData.On("GetAllFilteredByStatus", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(v.Output.Result, nil)

			//execute
			service := New(menteeData)
			_, err := service.GetAll(v.Input.Page, v.Input.Limit, v.Input.Status)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetOne(t *testing.T) {
	table := GetOneTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock data
			menteeData := new(mocks.MenteeData_)
			menteeData.On("GetOne", mock.AnythingOfType("int")).Return(v.Output.Result, nil)

			//execute
			service := New(menteeData)
			_, err := service.GetOne(v.Input.ID)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
