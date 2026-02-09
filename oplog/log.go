package oplog

import (
	"context"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/webomindapps-dev/coolaid-backend/config"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

/*
|--------------------------------------------------------------------------
| Context keys
|--------------------------------------------------------------------------
*/

type ctxKey string

const (
	requestIDKey ctxKey = "request_id"
	userIDKey    ctxKey = "user_id"
)

/*
|--------------------------------------------------------------------------
| Context helpers (OPTIONAL to use)
|--------------------------------------------------------------------------
*/

func AttachRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}

func AttachUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func contextFields(ctx context.Context) logrus.Fields {
	fields := logrus.Fields{
		"app": config.App.Name,
	}

	if ctx == nil {
		return fields
	}

	if reqID, ok := ctx.Value(requestIDKey).(string); ok && reqID != "" {
		fields["request_id"] = reqID
	}

	if userID, ok := ctx.Value(userIDKey).(string); ok && userID != "" {
		fields["user_id"] = userID
	}

	return fields
}

func callerName() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}
	if fn := runtime.FuncForPC(pc); fn != nil {
		return fn.Name()
	}
	return "unknown"
}

/*
|--------------------------------------------------------------------------
| Public logging functions
|--------------------------------------------------------------------------
*/

func Info(ctx context.Context, args ...interface{}) {
	fields := contextFields(ctx)
	fields["caller"] = callerName()

	logger.WithFields(fields).Info(args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	fields := contextFields(ctx)
	fields["caller"] = callerName()

	logger.WithFields(fields).Warn(args...)
}

func Error(ctx context.Context, args ...interface{}) {
	fields := contextFields(ctx)
	fields["caller"] = callerName()

	logger.WithFields(fields).Error(args...)
}
