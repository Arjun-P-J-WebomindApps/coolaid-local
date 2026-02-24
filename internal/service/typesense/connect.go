package typesense

import (
	"context"
	"errors"
	"time"

	"github.com/typesense/typesense-go/typesense"
	"github.com/webomindapps-dev/coolaid-backend/oplog"
)

type Context struct {
	Client *typesense.Client
}

func Connect(ctx context.Context, endpoint, apiKey string) (*Context, error) {

	if endpoint == "" {
		return nil, errors.New("typesense endpoint not configured")
	}

	oplog.Info(ctx, "connecting to typesense")

	client := typesense.NewClient(
		typesense.WithServer(endpoint),
		typesense.WithAPIKey(apiKey),
	)

	// Health check
	healthCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if _, err := client.Health(healthCtx, 5*time.Second); err != nil {
		oplog.Error(ctx, "typesense health check failed", "err=", err)
		return nil, err
	}

	oplog.Info(ctx, "typesense connected successfully")

	return &Context{
		Client: client,
	}, nil
}
