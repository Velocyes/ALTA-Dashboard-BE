package delivery

import (
	"alta-dashboard-be/features/logs"
)

func requestToEntity(logRequest logs.LogRequest) logs.LogEntity {
	return logs.LogEntity{
		Id:       uint(logRequest.Id),
		Title:    logRequest.Title,
		Status:   logRequest.Status,
		Feedback: logRequest.Feedback,
		UserID:   logRequest.UserID,
		MenteeID: logRequest.MenteeID,
	}
}

func entityToResponse(logEntity logs.LogEntity) logs.LogResponse {
	return logs.LogResponse{
		Id:        logEntity.Id,
		Title:     logEntity.Title,
		Status:    logEntity.Status,
		Feedback:  logEntity.Feedback,
		UserID:    logEntity.UserID,
		MenteeID:  logEntity.MenteeID,
		CreatedAt: logEntity.CreatedAt,
	}
}

func listEntityToResponse(logEntities []logs.LogEntity) []logs.LogResponse {
	var logResponses []logs.LogResponse
	for _, v := range logEntities {
		logResponses = append(logResponses, entityToResponse(v))
	}
	return logResponses
}
