package delivery

import (
	"alta-dashboard-be/features/class"
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
	tname := "test class update "
	return []TestTable{
		{
			Name: tname + "expect success",
			Input: struct {
				ID    interface{}
				Page  interface{}
				Limit interface{}
				Class []byte
			}{
				ID:    int(1),
				Class: []byte(`{"name":"matematika","short_name":"mtk 12","start_date":"2023-03-21","end_date":"2023-06-21"}`),
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: false,
			},
		},
		{
			Name: tname + "invalid start_date",
			Input: struct {
				ID    interface{}
				Page  interface{}
				Limit interface{}
				Class []byte
			}{
				ID:    int(1),
				Class: []byte(`{"name":"matematika","short_name":"mtk 12","start_date":"sadsa","end_date":"2023-06-21"}`),
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + "invalid end_date",
			Input: struct {
				ID    interface{}
				Page  interface{}
				Limit interface{}
				Class []byte
			}{
				ID:    int(1),
				Class: []byte(`{"name":"matematika","short_name":"mtk 12","start_date":"2023-03-21","end_date":"asdad"}`),
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + "invalid id",
			Input: struct {
				ID    interface{}
				Page  interface{}
				Limit interface{}
				Class []byte
			}{
				ID:    "asds",
				Class: []byte(`{"name":"matematika","short_name":"mtk 12","start_date":"2023-03-21","end_date":"2023-06-21"}`),
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
	tname := "test class create "
	return []TestTable{
		{
			Name: tname + "expect success",
			Input: struct {
				ID    interface{}
				Page  interface{}
				Limit interface{}
				Class []byte
			}{
				Class: []byte(`{"name":"matematika","short_name":"mtk 12","start_date":"2023-03-21","end_date":"2023-06-21"}`),
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: false,
			},
		},
		{
			Name: tname + "invalid start_date",
			Input: struct {
				ID    interface{}
				Page  interface{}
				Limit interface{}
				Class []byte
			}{
				Class: []byte(`{"name":"matematika","short_name":"mtk 12","start_date":"sadsa","end_date":"2023-06-21"}`),
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: true,
			},
		},
		{
			Name: tname + "invalid end_date",
			Input: struct {
				ID    interface{}
				Page  interface{}
				Limit interface{}
				Class []byte
			}{
				Class: []byte(`{"name":"matematika","short_name":"mtk 12","start_date":"2023-03-21","end_date":"asdad"}`),
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
	tname := "test class get one "
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
				Result: class.ClassCore{
					ID:        1,
					CreatedAt: time.Now(),
					Name:      "matematika",
					ShortName: "mtk 1",
					StartDate: time.Now().Add(24 * 1 * time.Hour),
					EndDate:   time.Now().Add(24 * 7 * time.Hour),
					UserID:    1,
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
				Result: class.ClassCore{
					ID:        1,
					CreatedAt: time.Now(),
					Name:      "matematika",
					ShortName: "mtk 1",
					StartDate: time.Now().Add(24 * 1 * time.Hour),
					EndDate:   time.Now().Add(24 * 7 * time.Hour),
					UserID:    1,
				},
			},
		},
	}
}

func DeleteTable() []TestTable {
	tname := "test class delete "
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
	tname := "test class get all "
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
				Result: []class.ClassCore{
					{
						ID:        1,
						CreatedAt: time.Now(),
						Name:      "matematika",
						ShortName: "mtk 1",
						StartDate: time.Now().Add(24 * 1 * time.Hour),
						EndDate:   time.Now().Add(24 * 7 * time.Hour),
						UserID:    1,
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
				Result: []class.ClassCore{
					{
						ID:        1,
						CreatedAt: time.Now(),
						Name:      "matematika",
						ShortName: "mtk 1",
						StartDate: time.Now().Add(24 * 1 * time.Hour),
						EndDate:   time.Now().Add(24 * 7 * time.Hour),
						UserID:    1,
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
				Result: []class.ClassCore{
					{
						ID:        1,
						CreatedAt: time.Now(),
						Name:      "matematika",
						ShortName: "mtk 1",
						StartDate: time.Now().Add(24 * 1 * time.Hour),
						EndDate:   time.Now().Add(24 * 7 * time.Hour),
						UserID:    1,
					},
				},
			},
		},
	}
}
