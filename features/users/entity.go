package users

import (
	"time"
)

type UserEntity struct {
	Id        uint
	FullName  string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string
	Team      string
	Role      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserServiceInterface interface {
	Login(email string, password string) (UserEntity, string, error)
	Create(input UserEntity, loggedInUserRole string) (UserEntity, error)
	GetAll(limit, offset int) (map[string]any, error)
	GetData(loggedInUserId, userId uint, loggedInUserRole string) (UserEntity, error)
	ModifyData(loggedInUserId, userId uint, loggedInUserRole string, input UserEntity) (UserEntity, error)
	Remove(loggedInUserId, userId uint, loggedInUserRole string) error
}

type UserDataInterface interface {
	Login(email string, password string) (UserEntity, string, error)
	Insert(input UserEntity) (UserEntity, error)
	SelectAll(limit, offset int) (map[string]any, error)
	SelectData(userId uint) (UserEntity, error)
	UpdateData(input UserEntity) (UserEntity, error)
	Delete(userId uint) error
}
