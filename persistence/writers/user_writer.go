package writers

import (
	"time"
	"user-api/persistence"
	"user-api/persistence/entities"

	"github.com/google/uuid"
)

type UserWriter struct{}

func (UserWriter) Create(user entities.User) *entities.User {
	db := persistence.NewConnection()
	defer db.Close()

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := db.NamedExec(
		`INSERT INTO "Users" ("Id", "Name", "Email", "CreatedAt", "UpdatedAt") 
		VALUES (:Id, :Name, :Email, :CreatedAt, :UpdatedAt)`,
		user,
	)
	if err != nil {
		return nil
	}

	return &user
}