package logs

import (
	"time"
)

type LogEntity struct {
	Id        uint
	Title     string `validate:"required"`
	Status    string `validate:"required"`
	Feedback  string `validate:"required"`
	UserID    uint   `validate:"required"`
	MenteeID  uint   `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type LogServiceInterface interface {
	Create(logInput LogEntity, loggedInUserId uint) (LogEntity, error)
	GetData(searchedMenteeId uint, limit, offset int) (map[string]any, error)
	// ModifyData(loggedInUserId, userId uint, loggedInUserRole string, input UserEntity) (logEntity, error)
	// Remove(loggedInUserId, userId uint, loggedInUserRole string) error
}

type LogDataInterface interface {
	Insert(input LogEntity) (LogEntity, error)
	SelectData(searchedMenteeId uint, limit, offset int) (map[string]any, error)
	// UpdateData(input UserEntity) (UserEntity, error)
	// Delete(userId uint) error
}
