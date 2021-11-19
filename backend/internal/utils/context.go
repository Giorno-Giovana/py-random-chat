package utils

import (
	"context"

	"junction-brella/internal/constants"
	"junction-brella/internal/model/core"
)

func CtxGetUserID(ctx context.Context) core.UserID {
	if value, ok := ctx.Value(constants.CtxKeyUserID{}).(string); ok {
		return core.UserID(value)
	}
	return ""
}
