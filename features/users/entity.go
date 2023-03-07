package users

import (
	"time"

	"github.com/labstack/echo/v4"
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

type UserRequest struct {
	Id       uint   `json:"id"`
	FullName string `json:"full_name" form:"full_name"`
	Email    string `json:"email" form:"email"`
	Team     string `json:"team" form:"team"`
	Role     string `json:"role" form:"role"`
	Status   string `json:"status" form:"status"`
}

type UserLogin struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserRegister struct {
	FullName string `json:"full_name" form:"full_name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Team     string `json:"team" form:"team"`
	Role     string `json:"role" form:"role"`
	Status   string `json:"status" form:"status"`
}

type UserResponse struct {
	Id       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Team     string `json:"team"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}

//go:generate mockery --name UserService_ --output ../../mocks
type UserServiceInterface_ interface {
	Login(email string, password string) (UserEntity, string, error)
	Create(input UserEntity, loggedInUserRole string) (UserEntity, error)
	GetAll(limit, offset int) (map[string]any, error)
	GetData(loggedInUserId, userId uint, loggedInUserRole string) (UserEntity, error)
	ModifyData(loggedInUserId, userId uint, loggedInUserRole string, input UserEntity) (UserEntity, error)
	Remove(loggedInUserId, userId uint, loggedInUserRole string) error
}

//go:generate mockery --name UserData_ --output ../../mocks
type UserDataInterface_ interface {
	Login(email string, password string) (UserEntity, string, error)
	Insert(input UserEntity) (UserEntity, error)
	SelectAll(limit, offset int) (map[string]any, error)
	SelectData(userId uint) (UserEntity, error)
	UpdateData(input UserEntity) (UserEntity, error)
	Delete(userId uint) error
}

//go:generate mockery --name UserDelivery_ --output ../../mocks
type UserDeliveryInterface_ interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
	GetAllUser(c echo.Context) error
	GetUserData(c echo.Context) error
	UpdateAccount(c echo.Context) error
	RemoveAccount(c echo.Context) error
}
