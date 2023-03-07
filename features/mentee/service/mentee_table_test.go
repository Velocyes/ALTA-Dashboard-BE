package service

import (
	"alta-dashboard-be/features/mentee"
	"time"
)

type TestTable struct {
	Name  string
	Input struct {
		ID     int
		UserID int
		Page   int
		Limit  int
		Mentee mentee.MenteeCore
	}
	Output struct {
		IsError bool
		Result  interface{}
	}
}

func pastTime() *time.Time {
	t, _ := time.Parse("2006-01-02", "2006-01-02")
	return &t
}

func GetOneTestTable() []TestTable {
	tname := "test get one mentee"
	return []TestTable{
		{
			Name: tname + " invalid id",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				ID: 0,
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
				Result: mentee.MenteeCore{

					ID:                1,
					CreatedAt:         time.Now(),
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationMajor:    "electrical engineering",
					EducationGradDate: pastTime(),
				},
			},
		},
		{
			Name: tname + " expect success",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				ID: 1,
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: false,
				Result: mentee.MenteeCore{

					ID:                1,
					CreatedAt:         time.Now(),
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationMajor:    "electrical engineering",
					EducationGradDate: pastTime(),
				},
			},
		},
	}
}

func GetAllTestTable() []TestTable {
	tname := "test get all mentee"
	return []TestTable{
		{
			Name: tname + " invalid page",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				Page:  0,
				Limit: 5,
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
				Result: []mentee.MenteeCore{
					{
						ID:                1,
						CreatedAt:         time.Now(),
						FullName:          "ahmad",
						Email:             "ahmad@ahmad.com",
						Address:           "jl ahmad no 15",
						Phone:             "088888888888",
						Telegram:          "ahmada123",
						EmergencyName:     "udin",
						EmergencyStatus:   "Keponakan",
						EmergencyPhone:    "088888888888",
						EducationType:     "IT",
						EducationMajor:    "electrical engineering",
						EducationGradDate: pastTime(),
					},
				},
			},
		},
		{
			Name: tname + " invalid limit",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				Page:  2,
				Limit: 0,
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
				Result: []mentee.MenteeCore{
					{
						ID:                1,
						CreatedAt:         time.Now(),
						FullName:          "ahmad",
						Email:             "ahmad@ahmad.com",
						Address:           "jl ahmad no 15",
						Phone:             "088888888888",
						Telegram:          "ahmada123",
						EmergencyName:     "udin",
						EmergencyStatus:   "Keponakan",
						EmergencyPhone:    "088888888888",
						EducationType:     "IT",
						EducationMajor:    "electrical engineering",
						EducationGradDate: pastTime(),
					},
				},
			},
		},
		{
			Name: tname + " expect success",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				Page:  1,
				Limit: 5,
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: false,
				Result: []mentee.MenteeCore{
					{
						ID:                1,
						CreatedAt:         time.Now(),
						FullName:          "ahmad",
						Email:             "ahmad@ahmad.com",
						Address:           "jl ahmad no 15",
						Phone:             "088888888888",
						Telegram:          "ahmada123",
						EmergencyName:     "udin",
						EmergencyStatus:   "Keponakan",
						EmergencyPhone:    "088888888888",
						EducationType:     "IT",
						EducationMajor:    "electrical engineering",
						EducationGradDate: pastTime(),
					},
				},
			},
		},
	}

}

func CreateTestTable() []TestTable {
	tname := "test create mentee"
	return []TestTable{
		{
			Name: tname + " null graduation date",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationMajor:    "electrical engineering",
					EducationGradDate: nil,
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: false,
			},
		},
		{
			Name: tname + " not null graduation date",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationGradDate: pastTime(),
					EducationMajor:    "electrical engineering",
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: false,
			},
		},
		{
			Name: tname + " invalid UserID",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				UserID: 0,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationGradDate: pastTime(),
					EducationMajor:    "electrical engineering",
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid full name",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad$#!~",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationMajor:    "electrical engineering",
					EducationGradDate: pastTime(),
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid email",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "x",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationMajor:    "electrical engineering",
					EducationGradDate: pastTime(),
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid address",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15$%^&*()",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationMajor:    "electrical engineering",
					EducationGradDate: pastTime(),
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid phone",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "0sad sad0$",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationGradDate: pastTime(),
					EducationMajor:    "electrical engineering",
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid telegram",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada 123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationGradDate: pastTime(),
					EducationMajor:    "electrical engineering",
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid emergency name",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin$%^&",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationMajor:    "electrical engineering",
					EducationGradDate: pastTime(),
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid emergency phone",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "08888 8",
					EducationType:     "IT",
					EducationMajor:    "electrical engineering",
					EducationGradDate: pastTime(),
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid emergency status",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "$%^&",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationMajor:    "electrical engineering",
					EducationGradDate: pastTime(),
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid education type",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "apaya",
					EducationMajor:    "electrical engineering",
					EducationGradDate: pastTime(),
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid education major",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationMajor:    "eledshufis13432#@1",
					EducationGradDate: pastTime(),
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
	}
}

func UpdateTestTable() []TestTable {
	tname := "test update mentee"
	return []TestTable{
		{
			Name: tname + " null graduation date",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				ID:     1,
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationMajor:    "electrical engineering",
					EducationGradDate: nil,
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: false,
			},
		},
		{
			Name: tname + " not null graduation date",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				ID:     1,
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationGradDate: pastTime(),
					EducationMajor:    "electrical engineering",
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: false,
			},
		},
		{
			Name: tname + " invalid UserID",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				ID:     1,
				UserID: 0,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationGradDate: pastTime(),
					EducationMajor:    "electrical engineering",
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid full name",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				ID:     1,
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad$#!~",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationMajor:    "electrical engineering",
					EducationGradDate: pastTime(),
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid email",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				ID:     1,
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "x",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationMajor:    "electrical engineering",
					EducationGradDate: pastTime(),
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid address",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				ID:     1,
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15$%^&*()",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationMajor:    "electrical engineering",
					EducationGradDate: pastTime(),
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid phone",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				ID:     1,
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "0sad sad0$",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationGradDate: pastTime(),
					EducationMajor:    "electrical engineering",
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid telegram",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				ID:     1,
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada 123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationGradDate: pastTime(),
					EducationMajor:    "electrical engineering",
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid emergency name",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				ID:     1,
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin$%^&",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationMajor:    "electrical engineering",
					EducationGradDate: pastTime(),
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid emergency phone",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				ID:     1,
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "08888 8",
					EducationType:     "IT",
					EducationMajor:    "electrical engineering",
					EducationGradDate: pastTime(),
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid emergency status",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				ID:     1,
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "$%^&",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationMajor:    "electrical engineering",
					EducationGradDate: pastTime(),
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid education type",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				ID:     1,
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "apaya",
					EducationMajor:    "electrical engineering",
					EducationGradDate: pastTime(),
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid education major",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				ID:     1,
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationMajor:    "eledshufis13432#@1",
					EducationGradDate: pastTime(),
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid ID",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				ID:     0,
				UserID: 2,
				Mentee: mentee.MenteeCore{
					FullName:          "ahmad",
					Email:             "ahmad@ahmad.com",
					Address:           "jl ahmad no 15",
					Phone:             "088888888888",
					Telegram:          "ahmada123",
					EmergencyName:     "udin",
					EmergencyStatus:   "Keponakan",
					EmergencyPhone:    "088888888888",
					EducationType:     "IT",
					EducationMajor:    "electrical engineering",
					EducationGradDate: pastTime(),
				},
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
	}
}

func DeleteTestTable() []TestTable {
	tname := "test delete mentee"
	return []TestTable{
		{
			Name: tname + " invalid UserID",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				UserID: 0,
				ID:     1,
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " invalid ID",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				UserID: 1,
				ID:     0,
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + " expect success",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Mentee mentee.MenteeCore
			}{
				UserID: 1,
				ID:     1,
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: false,
			},
		},
	}
}
