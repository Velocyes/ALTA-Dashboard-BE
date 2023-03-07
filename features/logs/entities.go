package logs

import (
	"time"

	"github.com/labstack/echo/v4"
)

type LogEntity struct {
	Id        uint
	Title     string `validate:"required"`
	Status    string 
	Feedback  string `validate:"required"`
	UserID    uint   `validate:"required"`
	MenteeID  uint   `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type LogRequest struct {
	Id       uint   `json:"id"`
	Title    string `json:"title" form:"title"`
	Status   string `json:"status" form:"status"`
	Feedback string `json:"feedback" form:"feedback"`
	UserID   uint   `json:"user_id" form:"user_id"`
	MenteeID uint   `json:"mentee_id" form:"mentee_id"`
}

type LogResponse struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	Feedback  string    `json:"feedback"`
	UserID    uint      `json:"user_id"`
	MenteeID  uint      `json:"mentee_id"`
	CreatedAt time.Time `json:"created_at"`
}

//go:generate mockery --name LogService_ --output ../../mocks
type LogServiceInterface_ interface {
	Create(logInput LogEntity, loggedInUserId uint) (LogEntity, error)
	GetData(searchedMenteeId uint, limit, offset int) (map[string]any, error)
}

//go:generate mockery --name type LogData_ --output ../../mocks
type LogDataInterface_ interface {
	Insert(input LogEntity) (LogEntity, error)
	SelectData(searchedMenteeId uint, limit, offset int) (map[string]any, error)
}

//go:generate mockery --name LogDelivery_ --output ../../mocks
type LogDeliveryInterface_ interface {
	AddLog(c echo.Context) error
	GetLogDataByMenteeId(c echo.Context) error
}
