package config

import (
	"context"

	"github.com/gin-gonic/gin"
	sql_models "github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
)

type ginContextKey struct{}
type userContextKey struct{}
type sessionContextKey struct{}

var (
	GinContextKey  = ginContextKey{}
	UserContextKey = userContextKey{}
	SessContextKey = sessionContextKey{}
)

func GinContextFromCtx(ctx context.Context) (*gin.Context, bool) {
	ginCtx, ok := ctx.Value(GinContextKey).(*gin.Context)
	return ginCtx, ok
}

func UserFromCtx(ctx context.Context) (*sql_models.User, bool) {
	user, ok := ctx.Value(UserContextKey).(*sql_models.User)
	return user, ok
}

func SessionFromCtx(ctx context.Context) (*sql_models.UserSession, bool) {
	sess, ok := ctx.Value(SessContextKey).(*sql_models.UserSession)
	return sess, ok
}
