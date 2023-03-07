package service

import (
	"alta-dashboard-be/features/class"
	"time"
)

type TestTable struct {
	Name  string
	Input struct {
		ID     int
		UserID int
		Page   int
		Limit  int
		Class  class.ClassCore
	}
	Output struct {
		IsError bool
		Result  interface{}
	}
}

func pastTime() time.Time {
	t, _ := time.Parse("2006-01-02", "2006-01-02")
	return t
}

func GetOneTestTable() []TestTable {
	tname := "test get one class"
	return []TestTable{
		{
			Name: tname + " invalid id",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Class  class.ClassCore
			}{
				ID: 0,
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
					ShortName: "mtk 2",
					StartDate: time.Now().Add(24 * 1 * time.Hour),
					EndDate:   time.Now().Add(24 * 7 * time.Hour),
					UserID:    1,
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
				Class  class.ClassCore
			}{
				ID: 2,
			},
			Output: struct {
				IsError bool
				Result  interface{}
			}{
				IsError: false,
				Result: class.ClassCore{
					ID:        2,
					CreatedAt: time.Now(),
					Name:      "matematika",
					ShortName: "mtk 2",
					StartDate: time.Now().Add(24 * 1 * time.Hour),
					EndDate:   time.Now().Add(24 * 7 * time.Hour),
					UserID:    1,
				},
			},
		},
	}
}

func GetAllTestTable() []TestTable {
	tname := "test get all class"
	return []TestTable{
		{
			Name: tname + " invalid page",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Class  class.ClassCore
			}{
				Page:  0,
				Limit: 3,
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
						ShortName: "mtk 2",
						StartDate: time.Now().Add(24 * 1 * time.Hour),
						EndDate:   time.Now().Add(24 * 7 * time.Hour),
						UserID:    1,
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
				Class  class.ClassCore
			}{
				Page:  2,
				Limit: -3,
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
						ShortName: "mtk 2",
						StartDate: time.Now().Add(24 * 1 * time.Hour),
						EndDate:   time.Now().Add(24 * 7 * time.Hour),
						UserID:    1,
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
				Class  class.ClassCore
			}{
				Page:  3,
				Limit: 3,
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
						ShortName: "mtk 2",
						StartDate: time.Now().Add(24 * 1 * time.Hour),
						EndDate:   time.Now().Add(24 * 7 * time.Hour),
						UserID:    1,
					},
					{
						ID:        2,
						CreatedAt: time.Now(),
						Name:      "fisika",
						ShortName: "fis 2",
						StartDate: time.Now().Add(24 * 1 * time.Hour),
						EndDate:   time.Now().Add(24 * 7 * time.Hour),
						UserID:    1,
					},
				},
			},
		},
	}

}

func CreateTestTable() []TestTable {
	tname := "test create class"
	return []TestTable{
		{
			Name: tname + "invalid UserID",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Class  class.ClassCore
			}{
				UserID: 0,
				Class: class.ClassCore{
					Name:      "matematika 23",
					ShortName: "mtk 23",
					StartDate: time.Now().Add(24 * 1 * time.Hour),
					EndDate:   time.Now().Add(24 * 7 * time.Hour),
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
			Name: tname + "invalid name",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Class  class.ClassCore
			}{
				UserID: 1,
				Class: class.ClassCore{
					Name:      "matematika @# %23",
					ShortName: "mtk 23",
					StartDate: time.Now().Add(24 * 1 * time.Hour),
					EndDate:   time.Now().Add(24 * 7 * time.Hour),
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
			Name: tname + "invalid short name",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Class  class.ClassCore
			}{
				UserID: 3,
				Class: class.ClassCore{
					Name:      "matematika 23",
					ShortName: "mtk @#_-23",
					StartDate: time.Now().Add(24 * 1 * time.Hour),
					EndDate:   time.Now().Add(24 * 7 * time.Hour),
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
			Name: tname + "invalid start date",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Class  class.ClassCore
			}{
				UserID: 2,
				Class: class.ClassCore{
					Name:      "matematika 23",
					ShortName: "mtk 23",
					StartDate: pastTime(),
					EndDate:   time.Now().Add(24 * 7 * time.Hour),
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
			Name: tname + "invalid end Date",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Class  class.ClassCore
			}{
				UserID: 0,
				Class: class.ClassCore{
					Name:      "matematika 23",
					ShortName: "mtk 23",
					StartDate: time.Now().Add(24 * 1 * time.Hour),
					EndDate:   pastTime(),
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
			Name: tname + "start date > end date",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Class  class.ClassCore
			}{
				UserID: 2,
				Class: class.ClassCore{
					Name:      "matematika 23",
					ShortName: "mtk 23",
					StartDate: time.Now().Add(24 * 7 * time.Hour),
					EndDate:   time.Now().Add(24 * 5 * time.Hour),
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
			Name: tname + "expect success",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Class  class.ClassCore
			}{
				UserID: 2,
				Class: class.ClassCore{
					Name:      "matematika 23",
					ShortName: "mtk 23",
					StartDate: time.Now().Add(24 * 1 * time.Hour),
					EndDate:   time.Now().Add(24 * 7 * time.Hour),
				},
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

func UpdateTestTable() []TestTable {
	tname := "test update class"
	return []TestTable{
		{
			Name: tname + "invalid UserID",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Class  class.ClassCore
			}{
				ID:     1,
				UserID: 0,
				Class: class.ClassCore{
					Name:      "matematika 23",
					ShortName: "mtk 23",
					StartDate: time.Now().Add(24 * 1 * time.Hour),
					EndDate:   time.Now().Add(24 * 7 * time.Hour),
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
			Name: tname + "invalid name",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Class  class.ClassCore
			}{
				ID:     2,
				UserID: 1,
				Class: class.ClassCore{
					Name:      "matematika @# %23",
					ShortName: "mtk 23",
					StartDate: time.Now().Add(24 * 1 * time.Hour),
					EndDate:   time.Now().Add(24 * 7 * time.Hour),
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
			Name: tname + "invalid short name",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Class  class.ClassCore
			}{
				ID:     2,
				UserID: 3,
				Class: class.ClassCore{
					Name:      "matematika 23",
					ShortName: "mtk @#_-23",
					StartDate: time.Now().Add(24 * 1 * time.Hour),
					EndDate:   time.Now().Add(24 * 7 * time.Hour),
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
			Name: tname + "invalid start date",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Class  class.ClassCore
			}{
				ID:     2,
				UserID: 2,
				Class: class.ClassCore{
					Name:      "matematika 23",
					ShortName: "mtk 23",
					StartDate: pastTime(),
					EndDate:   time.Now().Add(24 * 7 * time.Hour),
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
			Name: tname + "invalid end Date",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Class  class.ClassCore
			}{
				ID:     2,
				UserID: 0,
				Class: class.ClassCore{
					Name:      "matematika 23",
					ShortName: "mtk 23",
					StartDate: time.Now().Add(24 * 1 * time.Hour),
					EndDate:   pastTime(),
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
			Name: tname + "start date > end date",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Class  class.ClassCore
			}{
				ID:     2,
				UserID: 2,
				Class: class.ClassCore{
					Name:      "matematika 23",
					ShortName: "mtk 23",
					StartDate: time.Now().Add(24 * 7 * time.Hour),
					EndDate:   time.Now().Add(24 * 5 * time.Hour),
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
			Name: tname + "invalid ID",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Class  class.ClassCore
			}{
				ID:     0,
				UserID: 2,
				Class: class.ClassCore{
					Name:      "matematika 23",
					ShortName: "mtk 23",
					StartDate: time.Now().Add(24 * 1 * time.Hour),
					EndDate:   time.Now().Add(24 * 7 * time.Hour),
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
			Name: tname + "expect success",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Class  class.ClassCore
			}{
				ID:     2,
				UserID: 2,
				Class: class.ClassCore{
					Name:      "matematika 23",
					ShortName: "mtk 23",
					StartDate: time.Now().Add(24 * 1 * time.Hour),
					EndDate:   time.Now().Add(24 * 7 * time.Hour),
				},
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

func DeleteTestTable() []TestTable {
	tname := "test delete class"
	return []TestTable{
		{
			Name: tname + " invalid userID",
			Input: struct {
				ID     int
				UserID int
				Page   int
				Limit  int
				Class  class.ClassCore
			}{
				ID:     2,
				UserID: 0,
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
				Class  class.ClassCore
			}{
				ID:     0,
				UserID: 2,
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
				Class  class.ClassCore
			}{
				ID:     2,
				UserID: 2,
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
