package service

import (
	"alta-dashboard-be/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	table := LoginTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock data
			userDataMock := new(mocks.UserData_)
			userDataMock.On("Login", v.Input.Email, v.Input.Password).Return(v.Output.Result, v.Output.Token, nil)

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

// func TestCreate(t *testing.T) {
// 	table := CreateTestTable()
// 	for _, v := range table {
// 		t.Run(v.Name, func(t *testing.T) {
// 			//mock data
// 			userDataMock := new(mocks.UserData_)
// 			userDataMock.On("Insert", v.Input.userInput).Return(v.Output.Result, nil)

// 			//execute
// 			userService := New(userDataMock)
// 			_, err := userService.Create(v.Input.userInput, v.Input.loggedInUserRole)
// 			if v.Output.IsError {
// 				assert.NotNil(t, err)
// 			} else {
// 				assert.Nil(t, err)
// 			}
// 		})
// 	}
// }

func TestGetData(t *testing.T) {
	table := GetDataTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock data
			userDataMock := new(mocks.UserData_)
			userDataMock.On("SelectData", v.Input.userId).Return(v.Output.Result, nil)

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

func TestRemove(t *testing.T) {
	table := RemoveTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock data
			userDataMock := new(mocks.UserData_)
			userDataMock.On("Delete", v.Input.userId).Return(nil)

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
