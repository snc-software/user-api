package contracts

import "github.com/google/uuid"

// @name UserResponse
type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}