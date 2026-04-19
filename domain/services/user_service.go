package services

import (
	domainerrors "user-api/domain/errors"
	"user-api/domain/models"
	"user-api/persistence/entities"
	"user-api/persistence/readers"
	"user-api/persistence/writers"

	"github.com/google/uuid"
)

type UserService struct{}

func (UserService) GetByID(id uuid.UUID) (models.UserModel, error) {
	user, err := readers.UserReader{}.GetByID(id)
	if err != nil {
		return models.UserModel{}, domainerrors.NotFound
	}

	return models.UserModel{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (UserService) Create(createRequest models.CreateUserModel) (models.UserModel, error) {
	user, err := writers.UserWriter{}.Create(entities.User{
		Name:  createRequest.Name,
		Email: createRequest.Email,
	})
	if err != nil {
		return models.UserModel{}, domainerrors.Internal
	}

	return models.UserModel{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}