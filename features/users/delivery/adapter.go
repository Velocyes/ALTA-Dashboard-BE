package delivery

import "alta-dashboard-be/features/users"

func registerToEntity(userRegister users.UserRegister) users.UserEntity {
	return users.UserEntity{
		FullName: userRegister.FullName,
		Email:    userRegister.Email,
		Password: userRegister.Password,
		Team:     userRegister.Team,
		Role:     userRegister.Role,
		Status:   userRegister.Status,
	}
}

func requestToEntity(userRequest users.UserRequest) users.UserEntity {
	return users.UserEntity{
		Id:       uint(userRequest.Id),
		FullName: userRequest.FullName,
		Email:    userRequest.Email,
		Team:     userRequest.Team,
		Role:     userRequest.Role,
		Status:   userRequest.Status,
	}
}

func entityToResponse(userEntity users.UserEntity) users.UserResponse {
	return users.UserResponse{
		Id:       userEntity.Id,
		FullName: userEntity.FullName,
		Email:    userEntity.Email,
		Team:     userEntity.Team,
		Role:     userEntity.Role,
		Status:   userEntity.Status,
	}
}

func listEntityToResponse(userEntities []users.UserEntity) []users.UserResponse {
	var userResponses []users.UserResponse
	for _, v := range userEntities {
		userResponses = append(userResponses, entityToResponse(v))
	}
	return userResponses
}