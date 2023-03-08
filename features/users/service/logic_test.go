package service

import (
	"alta-dashboard-be/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	table := LoginTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock data
			userDataMock := new(mocks.UserData_)
			userDataMock.On("Login", mock.Anything, mock.AnythingOfType("string")).Return(v.Output.Result, v.Output.Token, nil)

			//execute
			userService := New(userDataMock)
			_, _, err := userService.Login(v.Input.Email, v.Input.Password)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	table := CreateTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock data
			userDataMock := new(mocks.UserData_)
			userDataMock.On("Insert", mock.Anything).Return(v.Output.Result, nil)

			//execute
			userService := New(userDataMock)
			_, err := userService.Create(v.Input.userInput, v.Input.loggedInUserRole)
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
			userDataMock := new(mocks.UserData_)
			userDataMock.On("SelectAll", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(v.Output.Result, nil)

			//execute
			userService := New(userDataMock)
			_, err := userService.GetAll(v.Input.limit, v.Input.offset)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetData(t *testing.T) {
	table := GetDataTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock data
			userDataMock := new(mocks.UserData_)
			userDataMock.On("SelectData", mock.Anything).Return(v.Output.Result, nil)

			//execute
			userService := New(userDataMock)
			_, err := userService.GetData(v.Input.loggedInUserId, v.Input.userId, v.Input.loggedInUserRole)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestModifyData(t *testing.T) {
	table := ModifyDataTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock data
			userDataMock := new(mocks.UserData_)
			userDataMock.On("UpdateData", mock.Anything).Return(v.Output.Result, nil)

			//execute
			userService := New(userDataMock)
			_, err := userService.ModifyData(v.Input.loggedInUserId, v.Input.userId, v.Input.loggedInUserRole, v.Input.userInput)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	table := RemoveTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock data
			userDataMock := new(mocks.UserData_)
			userDataMock.On("Delete", mock.AnythingOfType("uint")).Return(nil)

			//execute
			userService := New(userDataMock)
			err := userService.Remove(v.Input.loggedInUserId, v.Input.userId, v.Input.loggedInUserRole)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
