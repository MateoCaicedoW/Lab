package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id" `
	FirstName string    `db:"first_name" fako:"first_name"`
	LastName  string    `db:"last_name" fako:"last_name"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (u User) SelectLabel() string {
	return u.FirstName
}

func (u User) SelectValue() interface{} {
	return u.ID
}
