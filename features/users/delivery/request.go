package delivery

import "immersive-dashboard-app/features/users"

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

func registerToEntity(userRegister UserRegister) users.UserEntity {
	return users.UserEntity{
		FullName: userRegister.FullName,
		Email:    userRegister.Email,
		Password: userRegister.Password,
		Team:     userRegister.Team,
		Role:     userRegister.Role,
		Status:   userRegister.Status,
	}
}

func requestToEntity(userRequest UserRequest) users.UserEntity {
	return users.UserEntity{
		Id:       uint(userRequest.Id),
		FullName: userRequest.FullName,
		Email:    userRequest.Email,
		Team:     userRequest.Team,
		Role:     userRequest.Role,
		Status:   userRequest.Status,
	}
}
