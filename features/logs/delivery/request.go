package delivery

import (
	"alta-dashboard-be/features/logs"
)

type LogRequest struct {
	Id       uint   `json:"id"`
	Title    string `json:"title" form:"title"`
	Status   string `json:"status" form:"status"`
	Feedback string `json:"feedback" form:"feedback"`
	UserID   uint   `json:"user_id" form:"user_id"`
	MenteeID uint   `json:"mentee_id" form:"mentee_id"`
}

func requestToEntity(logRequest LogRequest) logs.LogEntity {
	return logs.LogEntity{
		Id:       uint(logRequest.Id),
		Title:    logRequest.Title,
		Status:   logRequest.Status,
		Feedback: logRequest.Feedback,
		UserID:   logRequest.UserID,
		MenteeID: logRequest.MenteeID,
	}
}
