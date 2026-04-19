package services

import (
	"user-api/domain/models"
	"user-api/persistence/entities"
	"user-api/persistence/readers"
	"user-api/persistence/writers"

	"github.com/google/uuid"
)

type UserService struct{}

func (UserService) GetByID(id uuid.UUID) *models.UserModel {
	user := readers.UserReader{}.GetByID(id)
	if user == nil {
		return nil
	}

	return &models.UserModel{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func (UserService) Create(createRequest models.CreateUserModel) *models.UserModel {
	user := writers.UserWriter{}.Create(entities.User{
		Name:  createRequest.Name,
		Email: createRequest.Email,
	})
	if user == nil {
		return nil
	}

	return &models.UserModel{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}