package lib

import (
	"github.com/gofrs/uuid"
)

// The purpose of this function is to generate a unique identifier string
func TokenGenerator() uuid.UUID {
	id, err := uuid.NewV4()
	if err != nil {
		return uuid.UUID{}
	}
	return id
}
