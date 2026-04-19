package services

import (
	"fmt"
	"user-api/domain/models"
	"user-api/exceptions"
	"user-api/persistence/entities"
	"user-api/persistence/readers"
	"user-api/persistence/writers"

	"github.com/google/uuid"
)

type UserService struct{}

func (UserService) GetByID(id uuid.UUID) (models.UserModel, error) {
	user, err := readers.UserReader{}.GetByID(id)
	if err != nil {

		return models.UserModel{}, exceptions.NotFound(fmt.Sprintf("User with id '%s' not found", id))
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
		return models.UserModel{}, exceptions.Internal()
	}

	return models.UserModel{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (UserService) GetPage(page, size int) ([]models.UserModel, int, error) {
	users, total, err := readers.UserReader{}.GetPage(page, size)
	if err != nil {
		return nil, 0, exceptions.Internal()
	}

	items := make([]models.UserModel, len(users))
	for i, user := range users {
		items[i] = models.UserModel{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
	}

	return items, total, nil
}

func (UserService) Delete(id uuid.UUID) error {
	_, err := readers.UserReader{}.GetByID(id)
	if err != nil {
		return exceptions.NotFound(fmt.Sprintf("User with id '%s' not found", id))
	}

	return writers.UserWriter{}.Delete(id)
}