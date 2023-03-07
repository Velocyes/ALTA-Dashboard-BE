package service

import (
	"alta-dashboard-be/features/logs"
	"alta-dashboard-be/mocks"
	"alta-dashboard-be/utils/consts"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	table := CreateTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock data
			logDataMock := new(mocks.LogData_)
			logDataMock.On("Insert", v.Input.logEntity).Return(v.Output.Result, nil)
			logDataMock.On("Insert", v.Input.logEntity).Return(v.Output.Result, errors.New(consts.SERVER_InternalServerError))

			//execute
			logService := New(logDataMock)
			logEntity, err := logService.Create(v.Input.logEntity, v.Input.LoggedInUserId)
			if v.Output.IsError {
				assert.Equal(t, logEntity, logs.LogEntity{})
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
			logDataMock := new(mocks.LogData_)
			logDataMock.On("SelectData", mock.AnythingOfType("uint"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(v.Output.Result, nil)

			//execute
			logService := New(logDataMock)
			_, err := logService.GetData(v.Input.logEntity.MenteeID, v.Input.Limit, v.Input.Offset)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
