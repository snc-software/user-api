package contracts

// @name CreateUserRequest
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}