package service

import (
	"alta-dashboard-be/features/logs"
	"alta-dashboard-be/utils/consts"
	"errors"
)

type TestTable struct {
	Name  string
	Input struct {
		LoggedInUserId   uint
		SearchedMenteeId uint
		Limit            int
		Offset           int
		logEntity        logs.LogEntity
	}
	Output struct {
		IsError   bool
		Result    interface{}
		errResult error
	}
}

func CreateTestTable() []TestTable {
	tname := "test create log"
	return []TestTable{
		{
			Name: tname + " invalid logged in user id",
			Input: struct {
				LoggedInUserId   uint
				SearchedMenteeId uint
				Limit            int
				Offset           int
				logEntity        logs.LogEntity
			}{
				LoggedInUserId: 0,
				logEntity:      logs.LogEntity{},
			},
			Output: struct {
				IsError   bool
				Result    interface{}
				errResult error
			}{
				IsError:   true,
				Result:    logs.LogEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " empty title",
			Input: struct {
				LoggedInUserId   uint
				SearchedMenteeId uint
				Limit            int
				Offset           int
				logEntity        logs.LogEntity
			}{
				LoggedInUserId: 1,
				logEntity: logs.LogEntity{
					Title:    "",
					Status:   consts.E_Log_None,
					Feedback: "Lorem Ipsum",
					MenteeID: 1,
				},
			},
			Output: struct {
				IsError   bool
				Result    interface{}
				errResult error
			}{
				IsError:   true,
				Result:    logs.LogEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " empty feedback",
			Input: struct {
				LoggedInUserId   uint
				SearchedMenteeId uint
				Limit            int
				Offset           int
				logEntity        logs.LogEntity
			}{
				LoggedInUserId: 1,
				logEntity: logs.LogEntity{
					Title:    "Lorem Ipsum",
					Status:   consts.E_Log_None,
					Feedback: "",
					MenteeID: 1,
				},
			},
			Output: struct {
				IsError   bool
				Result    interface{}
				errResult error
			}{
				IsError:   true,
				Result:    logs.LogEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " empty mentee id",
			Input: struct {
				LoggedInUserId   uint
				SearchedMenteeId uint
				Limit            int
				Offset           int
				logEntity        logs.LogEntity
			}{
				LoggedInUserId: 1,
				logEntity: logs.LogEntity{
					Title:    "Lorem Ipsum",
					Status:   consts.E_Log_None,
					Feedback: "Lorem Ipsum",
					MenteeID: 0,
				},
			},
			Output: struct {
				IsError   bool
				Result    interface{}
				errResult error
			}{
				IsError:   true,
				Result:    logs.LogEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " error response from data layer",
			Input: struct {
				LoggedInUserId   uint
				SearchedMenteeId uint
				Limit            int
				Offset           int
				logEntity        logs.LogEntity
			}{
				LoggedInUserId: 1,
				logEntity: logs.LogEntity{
					Title:    "Lorem Ipsum",
					Status:   consts.E_Log_None,
					Feedback: "Lorem Ipsum",
					MenteeID: 0,
				},
			},
			Output: struct {
				IsError   bool
				Result    interface{}
				errResult error
			}{
				IsError:   true,
				Result:    logs.LogEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " expect failed",
			Input: struct {
				LoggedInUserId   uint
				SearchedMenteeId uint
				Limit            int
				Offset           int
				logEntity        logs.LogEntity
			}{
				LoggedInUserId: 1,
				logEntity: logs.LogEntity{
					Title:    "Lorem Ipsum",
					Status:   consts.E_Log_None,
					Feedback: "Lorem Ipsum",
					UserID:   1,
					MenteeID: 1,
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
				errResult error
			}{
				IsError: false,
				Result: logs.LogEntity{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " expect success",
			Input: struct {
				LoggedInUserId   uint
				SearchedMenteeId uint
				Limit            int
				Offset           int
				logEntity        logs.LogEntity
			}{
				LoggedInUserId: 1,
				logEntity: logs.LogEntity{
					Title:    "Lorem Ipsum",
					Status:   consts.E_Log_None,
					Feedback: "Lorem Ipsum",
					UserID:   1,
					MenteeID: 1,
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
				errResult error
			}{
				IsError: false,
				Result: logs.LogEntity{
					Title:    "Lorem Ipsum",
					Status:   consts.E_Log_None,
					Feedback: "Lorem Ipsum",
					UserID:   1,
					MenteeID: 1,
				},
				errResult: nil,
			},
		},
	}
}

func GetDataTestTable() []TestTable {
	tname := "test get log"
	return []TestTable{
		{
			Name: tname + " expect failed",
			Input: struct {
				LoggedInUserId   uint
				SearchedMenteeId uint
				Limit            int
				Offset           int
				logEntity        logs.LogEntity
			}{
				SearchedMenteeId: 1,
				Limit:            1,
				Offset:           1,
			},
			Output: struct {
				IsError bool
				Result  interface{}
				errResult error
			}{
				IsError: true,
				Result:  map[string]any{},
				errResult: errors.New(""),
			},
		},
		{
			Name: tname + " expect success",
			Input: struct {
				LoggedInUserId   uint
				SearchedMenteeId uint
				Limit            int
				Offset           int
				logEntity        logs.LogEntity
			}{
				SearchedMenteeId: 1,
				Limit:            1,
				Offset:           1,
			},
			Output: struct {
				IsError bool
				Result  interface{}
				errResult error
			}{
				IsError: false,
				Result:  map[string]any{},
				errResult: nil,
			},
		},
	}
}
