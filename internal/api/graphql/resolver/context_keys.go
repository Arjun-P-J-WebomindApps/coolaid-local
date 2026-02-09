package resolvers

import (
	"context"

	"github.com/webomindapps-dev/coolaid-backend/internal/domain/auth"
)

type CookieWriterKey struct{}

func cookieWriterFromContext(
	ctx context.Context,
) (auth.CookieWriter, bool) {

	writer, ok := ctx.Value(CookieWriterKey{}).(auth.CookieWriter)
	return writer, ok
}
