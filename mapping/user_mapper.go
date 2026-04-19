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

func MapToPagedResponse(users []models.UserModel, page, size, total int) contracts.PagedResponse[contracts.UserResponse] {
	items := make([]contracts.UserResponse, len(users))
	for i, user := range users {
		items[i] = MapToResponse(user)
	}

	return contracts.PagedResponse[contracts.UserResponse]{
		Items: items,
		Pagination: contracts.Pagination{
			Page:  page,
			Size:  size,
			Total: total,
		},
	}
}