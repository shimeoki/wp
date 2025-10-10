package domain

import "github.com/google/uuid"

type ID uuid.UUID

func NewID() ID {
	return ID(uuid.Must(uuid.NewV7()))
}

func ParseID(value string) (ID, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return ID(uuid.Nil), err
	}

	return ID(id), nil
}

func (id ID) String() string {
	return (uuid.UUID)(id).String()
}
