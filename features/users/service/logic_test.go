package service

// import (
// 	"clean-arch/features/users"
// 	"clean-arch/mocks"
// 	"errors"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestCreate(t *testing.T) {
// 	userDataMock := new(mocks.UserDataMock)
// 	inputData := users.UserEntity{Id: 1, Name: "Joko", Email: "Joko@gmail.com", Password: "qwerty", Address: "Jakarta", Role: "User"}

// 	t.Run("Failed validate", func(t *testing.T) {
// 		inputDataCopy := inputData
// 		inputDataCopy.Email = ""
// 		userService := New(userDataMock)
// 		err := userService.Create(inputDataCopy)
// 		assert.NotNil(t, err)
// 		userDataMock.AssertExpectations(t)
// 	})

// 	t.Run("Failed Insert User Error", func(t *testing.T) {
// 		userDataMock.On("Insert", inputData).Return(errors.New("error insert data")).Once()

// 		userService := New(userDataMock)
// 		err := userService.Create(inputData)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, "error insert data", err.Error())
// 		userDataMock.AssertExpectations(t)
// 	})

// 	t.Run("Success Insert User", func(t *testing.T) {
// 		userDataMock.On("Insert", inputData).Return(nil).Once()

// 		userService := New(userDataMock)
// 		err := userService.Create(inputData)
// 		assert.Nil(t, err)
// 		userDataMock.AssertExpectations(t)
// 	})
// }

// func TestGetAll(t *testing.T) {
// 	userDataMock := new(mocks.UserDataMock)
// 	returnData := []users.UserEntity{
// 		{Id: 1, Name: "Joko", Email: "Joko@gmail.com", Password: "qwerty", Address: "Jakarta", Role: "User"},
// 	}

// 	t.Run("Success Get All Users", func(t *testing.T) {
// 		userDataMock.On("SelectAll").Return(returnData, nil).Once()

// 		userService := New(userDataMock)
// 		response, err := userService.GetAll()
// 		assert.Nil(t, err)
// 		assert.Equal(t, returnData[0].Name, response[0].Name)
// 		userDataMock.AssertExpectations(t)
// 	})
// }

// func TestGetUser(t *testing.T) {
// 	userDataMock := new(mocks.UserDataMock)
// 	userId := uint(1)
// 	returnData := users.UserEntity{
// 		Id: 1, Name: "Joko", Email: "Joko@gmail.com", Password: "qwerty", Address: "Jakarta", Role: "User",
// 	}

// 	t.Run("Success Get User", func(t *testing.T) {
// 		userDataMock.On("SelectByUserId", userId).Return(returnData, nil).Once()

// 		userService := New(userDataMock)
// 		response, err := userService.GetUser(userId)
// 		assert.Nil(t, err)
// 		assert.Equal(t, returnData.Name, response.Name)
// 		userDataMock.AssertExpectations(t)
// 	})
// }

// func TestModify(t *testing.T) {
// 	userDataMock := new(mocks.UserDataMock)
// 	userId := 1
// 	inputData := users.UserEntity{
// 		Id: 1, Name: "Joko", Email: "Joko@gmail.com", Password: "qwerty", Address: "Jakarta", Role: "User",
// 	}

// 	t.Run("Success Update User", func(t *testing.T) {
// 		userDataMock.On("Update", uint(userId), inputData).Return(nil).Once()

// 		userService := New(userDataMock)
// 		err := userService.Modify(uint(userId), inputData)
// 		assert.Nil(t, err)
// 		userDataMock.AssertExpectations(t)
// 	})

// 	t.Run("Failed Find User By ID", func(t *testing.T) {
// 		userDataMock.On("Update", uint(userId), inputData).Return(errors.New("error select user")).Once()

// 		userService := New(userDataMock)
// 		err := userService.Modify(uint(userId), inputData)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, "error select user", err.Error())
// 		userDataMock.AssertExpectations(t)
// 	})

// 	t.Run("Failed Update User", func(t *testing.T) {
// 		userDataMock.On("Update", uint(userId), inputData).Return(errors.New("update error, row affected = 0")).Once()

// 		userService := New(userDataMock)
// 		err := userService.Modify(uint(userId), inputData)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, "update error, row affected = 0", err.Error())
// 		userDataMock.AssertExpectations(t)
// 	})
// }

// func TestRemove(t *testing.T) {
// 	userDataMock := new(mocks.UserDataMock)
// 	userId := 1

// 	t.Run("Success Delete User", func(t *testing.T) {
// 		userDataMock.On("Delete", uint(userId)).Return(nil).Once()

// 		userService := New(userDataMock)
// 		err := userService.Remove(uint(userId))
// 		assert.Nil(t, err)
// 		userDataMock.AssertExpectations(t)
// 	})

// 	t.Run("Failed Find User By ID", func(t *testing.T) {
// 		userDataMock.On("Delete", uint(userId)).Return(errors.New("delete error, row affected = 0")).Once()

// 		userService := New(userDataMock)
// 		err := userService.Remove(uint(userId))
// 		assert.NotNil(t, err)
// 		assert.Equal(t, "delete error, row affected = 0", err.Error())
// 		userDataMock.AssertExpectations(t)
// 	})
// }
