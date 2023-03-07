package delivery

import (
	"alta-dashboard-be/features/mentee"
	"time"
)

type TestTable struct {
	Name  string
	Input struct {
		ID    interface{}
		Page  interface{}
		Limit interface{}
		Class []byte
	}
	Output struct {
		IsError bool
		Result  interface{}
	}
}

func UpdateTable() []TestTable {
	tname := "test mentee update "
	return []TestTable{
		{
			Name: tname + "expect success grad date null",
			Input: struct {
				ID    interface{}
				Page  interface{}
				Limit interface{}
				Class []byte
			}{
				ID:    int(1),
				Class: []byte(`{"fullname":"ahmad","email":"ahmad@ahmad.com","address":"jl batu","phone":"081218666666","telegram":"itsahmad","emergency_name":"abdul","emergency_phone":"081218000000","emergency_status":"saudara","education_type":"IT","education_major":"teknik elektro","class_id":1}`),
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: false,
			},
		},
		{
			Name: tname + "expect success grad date not null",
			Input: struct {
				ID    interface{}
				Page  interface{}
				Limit interface{}
				Class []byte
			}{
				ID:    int(1),
				Class: []byte(`{"fullname":"ahmad","email":"ahmad@ahmad.com","address":"jl batu","phone":"081218666666","telegram":"itsahmad","emergency_name":"abdul","emergency_phone":"081218000000","emergency_status":"saudara","education_type":"IT","education_major":"teknik elektro","education_grad_date":"2023-01-20","class_id":1}`),
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: false,
			},
		},
		{
			Name: tname + "invalid ID",
			Input: struct {
				ID    interface{}
				Page  interface{}
				Limit interface{}
				Class []byte
			}{
				ID:    "sadad",
				Class: []byte(`{"fullname":"ahmad","email":"ahmad@ahmad.com","address":"jl batu","phone":"081218666666","telegram":"itsahmad","emergency_name":"abdul","emergency_phone":"081218000000","emergency_status":"saudara","education_type":"IT","education_major":"teknik elektro","education_grad_date":"2023-01-20","class_id":1}`),
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + "invalid graduate date",
			Input: struct {
				ID    interface{}
				Page  interface{}
				Limit interface{}
				Class []byte
			}{
				ID:    1,
				Class: []byte(`{"fullname":"ahmad","email":"ahmad@ahmad.com","address":"jl batu","phone":"081218666666","telegram":"itsahmad","emergency_name":"abdul","emergency_phone":"081218000000","emergency_status":"saudara","education_type":"IT","education_major":"teknik elektro","education_grad_date":"sdfasd","class_id":1}`),
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

func CreateTable() []TestTable {
	tname := "test mentee create "
	return []TestTable{
		{
			Name: tname + "expect success grad date null",
			Input: struct {
				ID    interface{}
				Page  interface{}
				Limit interface{}
				Class []byte
			}{
				Class: []byte(`{"fullname":"ahmad","email":"ahmad@ahmad.com","address":"jl batu","phone":"081218666666","telegram":"itsahmad","emergency_name":"abdul","emergency_phone":"081218000000","emergency_status":"saudara","education_type":"IT","education_major":"teknik elektro","class_id":1}`),
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: false,
			},
		},
		{
			Name: tname + "expect success grad date not null",
			Input: struct {
				ID    interface{}
				Page  interface{}
				Limit interface{}
				Class []byte
			}{
				Class: []byte(`{"fullname":"ahmad","email":"ahmad@ahmad.com","address":"jl batu","phone":"081218666666","telegram":"itsahmad","emergency_name":"abdul","emergency_phone":"081218000000","emergency_status":"saudara","education_type":"IT","education_major":"teknik elektro","education_grad_date":"2023-01-20","class_id":1}`),
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: false,
			},
		},
		{
			Name: tname + "invalid graduate date",
			Input: struct {
				ID    interface{}
				Page  interface{}
				Limit interface{}
				Class []byte
			}{
				Class: []byte(`{"fullname":"ahmad","email":"ahmad@ahmad.com","address":"jl batu","phone":"081218666666","telegram":"itsahmad","emergency_name":"abdul","emergency_phone":"081218000000","emergency_status":"saudara","education_type":"IT","education_major":"teknik elektro","education_grad_date":"sdfasd","class_id":1}`),
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

func GetOneTable() []TestTable {
	tname := "test mentee get one "
	return []TestTable{
		{
			Name: tname + "invalid ID",
			Input: struct {
				ID    interface{}
				Page  interface{}
				Limit interface{}
				Class []byte
			}{
				ID: "asd",
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
					Address:           "jl batu",
					Phone:             "0812111111",
					Telegram:          "itsahmad",
					EmergencyName:     "abdul",
					EmergencyPhone:    "08121110000",
					EmergencyStatus:   "keponakan",
					EducationType:     "IT",
					EducationMajor:    "teknik elektro",
					EducationGradDate: &time.Time{},
					ClassID:           1,
				},
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				ID    interface{}
				Page  interface{}
				Limit interface{}
				Class []byte
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
					Address:           "jl batu",
					Phone:             "0812111111",
					Telegram:          "itsahmad",
					EmergencyName:     "abdul",
					EmergencyPhone:    "08121110000",
					EmergencyStatus:   "keponakan",
					EducationType:     "IT",
					EducationMajor:    "teknik elektro",
					EducationGradDate: &time.Time{},
					ClassID:           1,
				},
			},
		},
	}
}

func DeleteTable() []TestTable {
	tname := "test mentee delete "
	return []TestTable{
		{
			Name: tname + "invalid ID",
			Input: struct {
				ID    interface{}
				Page  interface{}
				Limit interface{}
				Class []byte
			}{
				ID: "asd",
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + "expect success",
			Input: struct {
				ID    interface{}
				Page  interface{}
				Limit interface{}
				Class []byte
			}{
				ID: 1,
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

func GetAllTable() []TestTable {
	tname := "test mentee get all "
	return []TestTable{
		{
			Name: tname + "expect success",
			Input: struct {
				ID    interface{}
				Page  interface{}
				Limit interface{}
				Class []byte
			}{
				Page:  3,
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
						Address:           "jl batu",
						Phone:             "0812111111",
						Telegram:          "itsahmad",
						EmergencyName:     "abdul",
						EmergencyPhone:    "08121110000",
						EmergencyStatus:   "keponakan",
						EducationType:     "IT",
						EducationMajor:    "teknik elektro",
						EducationGradDate: &time.Time{},
						ClassID:           1,
					},
				},
			},
		},
		{
			Name: tname + "invalid page",
			Input: struct {
				ID    interface{}
				Page  interface{}
				Limit interface{}
				Class []byte
			}{
				Page:  "sada",
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
						Address:           "jl batu",
						Phone:             "0812111111",
						Telegram:          "itsahmad",
						EmergencyName:     "abdul",
						EmergencyPhone:    "08121110000",
						EmergencyStatus:   "keponakan",
						EducationType:     "IT",
						EducationMajor:    "teknik elektro",
						EducationGradDate: &time.Time{},
						ClassID:           1,
					},
				},
			},
		},
		{
			Name: tname + "invalid limit",
			Input: struct {
				ID    interface{}
				Page  interface{}
				Limit interface{}
				Class []byte
			}{
				Page:  3,
				Limit: "asd",
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
						Address:           "jl batu",
						Phone:             "0812111111",
						Telegram:          "itsahmad",
						EmergencyName:     "abdul",
						EmergencyPhone:    "08121110000",
						EmergencyStatus:   "keponakan",
						EducationType:     "IT",
						EducationMajor:    "teknik elektro",
						EducationGradDate: &time.Time{},
						ClassID:           1,
					},
				},
			},
		},
	}
}
