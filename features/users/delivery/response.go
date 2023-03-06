package delivery

import (
	"immersive-dashboard-app/features/users"
)

type UserResponse struct {
	Id       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Team     string `json:"team"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}

func entityToResponse(userEntity users.UserEntity) UserResponse {
	return UserResponse{
		Id:       userEntity.Id,
		FullName: userEntity.FullName,
		Email:    userEntity.Email,
		Team:     userEntity.Team,
		Role:     userEntity.Role,
		Status:   userEntity.Status,
	}
}

func listEntityToResponse(userEntities []users.UserEntity) []UserResponse {
	var userResponses []UserResponse
	for _, v := range userEntities {
		userResponses = append(userResponses, entityToResponse(v))
	}
	return userResponses
}
