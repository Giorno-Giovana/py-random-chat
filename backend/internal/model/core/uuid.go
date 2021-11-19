// This file is in core bc GenUUID generates
// UUID that are needed for core models
package core

import (
	"fmt"

	"github.com/gofrs/uuid"

	"junction-brella/internal/constants"
)

func GenUUID() (string, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("%w: %v", constants.ErrGenerateUUID, err)
	}
	return uuid.String(), nil
}
