package writers

import (
	"fmt"
	"time"
	"user-api/exceptions"
	"user-api/persistence"
	"user-api/persistence/entities"

	"github.com/google/uuid"
)

type UserWriter struct{}

func (UserWriter) Create(user entities.User) (entities.User, error) {
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
		return entities.User{}, err
	}

	return user, nil
}

func (UserWriter) Delete(id uuid.UUID) error {
	db := persistence.NewConnection()
	defer db.Close()

	result, err := db.Exec(`DELETE FROM "Users" WHERE "Id" = $1`, id)
	if err != nil {
		return exceptions.Internal()
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return exceptions.NotFound(fmt.Sprintf("User with id '%s' not found", id))
	}

	return nil
}