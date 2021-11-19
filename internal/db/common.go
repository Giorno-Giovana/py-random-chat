package db

import (
	"errors"

	"gitlab.com/procsy/attorney/api/internal/constants"
	"go.mongodb.org/mongo-driver/mongo"
)

func wrapError(err error) error {
	if errors.Is(err, mongo.ErrNoDocuments) {
		return constants.ErrDBNotFound
	}
	return err
}
