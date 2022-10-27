package models

import (
	"github.com/gofrs/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id" `
	FirstName string    `db:"first_name" fako:"first_name"`
	LastName  string    `db:"last_name" fako:"last_name"`
}
