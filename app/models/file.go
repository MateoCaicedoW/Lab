package models

import (
	"github.com/gobuffalo/buffalo/binding"
	"github.com/gofrs/uuid"
)

type File struct {
	UserID uuid.UUID `db:"-"`
	// File is a pointer to a binding.File
	MyFile binding.File `db:"-"`
}
