package log

import (
	"context"

	"github.com/sirupsen/logrus"
	"gitlab.com/procsy/attorney/api/internal/constants"
)

func WithFieldsFromContext(ctx context.Context, log *logrus.Entry) *logrus.Entry {
	if value, ok := ctx.Value(constants.CtxKeyUserID{}).(string); ok {
		log = log.WithField("user_id", value)
	}

	if value, ok := ctx.Value(constants.CtxKeyXRequestID{}).(string); ok {
		log = log.WithField("x_request_id", value)
	}

	return log
}
