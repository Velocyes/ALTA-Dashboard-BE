package mentee

import (
	"time"

	"github.com/labstack/echo/v4"
)

type MenteeCore struct {
	ID                int
	CreatedAt         time.Time
	FullName          string
	Email             string
	Address           string
	Phone             string
	Telegram          string
	EmergencyName     string
	EmergencyPhone    string
	EmergencyStatus   string
	EducationType     string
	EducationMajor    string
	EducationGradDate *time.Time //nullable
	ClassID           int
}

type MenteeRequest struct {
	FullName          string
	Email             string
	Address           string
	Phone             string
	Telegram          string
	EmergencyName     string
	EmergencyPhone    string
	EmergencyStatus   string
	EducationType     string
	EducationMajor    string
	EducationGradDate string
	ClassID           int
}

type MenteeResponse struct {
	ID                int
	CreatedAt         time.Time
	FullName          string
	Email             string
	Address           string
	Phone             string
	Telegram          string
	EmergencyName     string
	EmergencyPhone    string
	EmergencyStatus   string
	EducationType     string
	EducationMajor    string
	EducationGradDate string
	ClassID           int
}

//go:generate mockery --name MenteeData_ --output ../../mocks
type MenteeData_ interface {
	Create(mentee MenteeCore) error
	GetAll(page int, limit int) ([]MenteeCore, error)
	GetOne(id int) (MenteeCore, error)
	Update(id int, mentee MenteeCore) error
	Delete(id int) error
}

//go:generate mockery --name MenteeService_ --output ../../mocks
type MenteeService_ interface {
	Create(userID int, mentee MenteeCore) error
	GetAll(page int, limit int) ([]MenteeCore, error)
	GetOne(id int) (MenteeCore, error)
	Update(userID int, id int, mentee MenteeCore) error
	Delete(userID int, id int) error
}

//go:generate mockery --name MenteeDelivery_ --output ../../mocks
type MenteeDelivery_ interface {
	Create(c echo.Context) error
	GetAll(c echo.Context) error
	GetOne(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}
