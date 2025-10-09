package domain

import "github.com/google/uuid"

type ID uuid.UUID

func NewID() ID {
	return ID(uuid.Must(uuid.NewV7()))
}

func (id ID) String() string {
	return (uuid.UUID)(id).String()
}
