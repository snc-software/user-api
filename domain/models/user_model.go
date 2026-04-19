package models

import "github.com/google/uuid"

type UserModel struct {
	ID    uuid.UUID
	Name  string
	Email string
}