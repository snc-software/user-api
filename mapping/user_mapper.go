package mapping

import (
	"user-api/domain/models"
	"user-api/routes/contracts"
)

func MapToDomain(createRequest contracts.CreateUserRequest) models.CreateUserModel {
	return models.CreateUserModel{
		Name:  createRequest.Name,
		Email: createRequest.Email,
	}
}

func MapToResponse(user models.UserModel) contracts.UserResponse {
	return contracts.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}
}