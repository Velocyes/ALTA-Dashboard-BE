package delivery

import (
	"alta-dashboard-be/features/logs"
	"time"
)

type LogResponse struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	Feedback  string    `json:"feedback"`
	UserID    uint      `json:"user_id"`
	MenteeID  uint      `json:"mentee_id"`
	CreatedAt time.Time `json:"created_at"`
}

func entityToResponse(logEntity logs.LogEntity) LogResponse {
	return LogResponse{
		Id:        logEntity.Id,
		Title:     logEntity.Title,
		Status:    logEntity.Status,
		Feedback:  logEntity.Feedback,
		UserID:    logEntity.UserID,
		MenteeID:  logEntity.MenteeID,
		CreatedAt: logEntity.CreatedAt,
	}
}

func listEntityToResponse(logEntities []logs.LogEntity) []LogResponse {
	var logResponses []LogResponse
	for _, v := range logEntities {
		logResponses = append(logResponses, entityToResponse(v))
	}
	return logResponses
}
