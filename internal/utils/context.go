package utils

import (
	"context"

	"gitlab.com/procsy/attorney/api/internal/constants"
	"gitlab.com/procsy/attorney/api/internal/model/core"
)

func CtxGetUserID(ctx context.Context) core.UserID {
	if value, ok := ctx.Value(constants.CtxKeyUserID{}).(string); ok {
		return core.UserID(value)
	}
	return ""
}
