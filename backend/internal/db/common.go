package db

import (
	"errors"

	"junction-brella/internal/constants"

	"go.mongodb.org/mongo-driver/mongo"
)

func wrapError(err error) error {
	if errors.Is(err, mongo.ErrNoDocuments) {
		return constants.ErrDBNotFound
	}
	return err
}
