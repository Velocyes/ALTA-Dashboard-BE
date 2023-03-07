package class

import (
	"time"

	"github.com/labstack/echo/v4"
)

type ClassCore struct {
	ID        int
	CreatedAt time.Time
	Name      string
	ShortName string
	StartDate time.Time
	EndDate   time.Time
	UserID    int
}

type ClassRequest struct {
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type ClassResponse struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	UserID    int    `json:"user_id"`
}

//go:generate mockery --name ClassData_ --output ../../mocks
type ClassData_ interface {
	Create(class ClassCore) error
	GetAll(page int, limit int) ([]ClassCore, error)
	GetOne(id int) (ClassCore, error)
	Update(id int, class ClassCore) error
	Delete(id int) error
}

//go:generate mockery --name ClassService_ --output ../../mocks
type ClassService_ interface {
	Create(userID int, class ClassCore) error
	GetAll(page int, limit int) ([]ClassCore, error)
	GetOne(id int) (ClassCore, error)
	Update(userID int, id int, class ClassCore) error
	Delete(userID int, id int) error
}

//go:generate mockery --name ClassDelivery_ --output ../../mocks

type EchoClassContext interface {
	Bind(interface{}) error
	Param(string) string
	QueryParam(string) string
	JSON(int, interface{}) error
}

type ClassDelivery_ interface {
	Create(c echo.Context) error
	GetAll(c echo.Context) error
	GetOne(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}
