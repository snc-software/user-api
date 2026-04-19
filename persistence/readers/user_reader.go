package readers

import (
	"user-api/persistence"
	"user-api/persistence/entities"

	"github.com/google/uuid"
)

type UserReader struct{}

func (UserReader) GetByID(id uuid.UUID) (entities.User, error) {
	db := persistence.NewConnection()
	defer db.Close()

	var user entities.User
	err := db.Get(&user, `SELECT "Id", "Name", "Email", "CreatedAt", "UpdatedAt" FROM "Users" WHERE "Id" = $1`, id)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (UserReader) GetPage(page, size int) ([]entities.User, int, error) {
	db := persistence.NewConnection()
	defer db.Close()

	var total int
	err := db.Get(&total, `SELECT COUNT(*) FROM "Users"`)
	if err != nil {
		return nil, 0, err
	}

	var users []entities.User
	offset := (page - 1) * size
	err = db.Select(&users, `SELECT "Id", "Name", "Email", "CreatedAt", "UpdatedAt" FROM "Users" ORDER BY "CreatedAt" DESC LIMIT $1 OFFSET $2`, size, offset)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
