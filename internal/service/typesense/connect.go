package typesense

import (
	"context"
	"errors"
	"time"

	"github.com/typesense/typesense-go/typesense"
	"github.com/webomindapps-dev/coolaid-backend/config"
	"github.com/webomindapps-dev/coolaid-backend/oplog"
)

type Context struct {
	Client *typesense.Client
}

func Connect(ctx context.Context) (*Context, error) {

	if config.SearchEngine.TypesenseAPIEndpoint == "" {
		return nil, errors.New("typesense endpoint not configured")
	}

	oplog.Info(ctx, "connecting to typesense")

	client := typesense.NewClient(
		typesense.WithServer(config.SearchEngine.TypesenseAPIEndpoint),
		typesense.WithAPIKey(config.SearchEngine.TypesenseAPIKey),
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
