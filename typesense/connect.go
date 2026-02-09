package typesense

import (
	"context"
	"errors"
	"time"

	"github.com/typesense/typesense-go/typesense"
	"github.com/webomindapps-dev/coolaid-backend/config"
	"github.com/webomindapps-dev/coolaid-backend/oplog"
)

type TypesenseContext struct {
	Client *typesense.Client
}

var TS *TypesenseContext

func Connect(ctx context.Context) error {
	if config.SearchEngine.TypesenseAPIEndpoint == "" {
		return errors.New("typesense endpoint not configured")
	}

	oplog.Info(ctx, "connecting to typesense")

	client := typesense.NewClient(
		typesense.WithServer(config.SearchEngine.TypesenseAPIEndpoint),
		typesense.WithAPIKey(config.SearchEngine.TypesenseAPIKey),
	)

	// ✅ Correct health check
	healthCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if _, err := client.Health(healthCtx, 5*time.Second); err != nil {
		oplog.Error(ctx, "typesense health check failed", err)
		return err
	}

	TS = &TypesenseContext{
		Client: client,
	}

	oplog.Info(ctx, "typesense connected successfully")
	return nil
}
