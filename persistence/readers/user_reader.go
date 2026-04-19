package readers

import (
	"user-api/persistence"
	"user-api/persistence/entities"

	"github.com/google/uuid"
)

type UserReader struct{}

func (UserReader) GetByID(id uuid.UUID) *entities.User {
	db := persistence.NewConnection()
	defer db.Close()

	var user entities.User
	err := db.Get(&user, `SELECT "Id", "Name", "Email", "CreatedAt", "UpdatedAt" FROM "Users" WHERE "Id" = $1`, id)
	if err != nil {
		return nil
	}

	return &user
}