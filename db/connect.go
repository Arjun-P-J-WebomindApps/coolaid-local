package db

import (
	"context"
	"database/sql"
	"strings"

	_ "github.com/lib/pq"

	"github.com/webomindapps-dev/coolaid-backend/config"
	sql_models "github.com/webomindapps-dev/coolaid-backend/internal/generated/sqlc"
	"github.com/webomindapps-dev/coolaid-backend/oplog"
)

type DBContext struct {
	SqlDB   *sql.DB
	Queries *sql_models.Queries
}

var DB *DBContext

func Connect(ctx context.Context) error {
	dsn := getPostgresDsn()

	oplog.Info(ctx, "connecting to database")

	database, err := sql.Open("postgres", dsn)
	if err != nil {
		oplog.Error(ctx, "failed to open database", err)
		return err
	}

	if err := database.Ping(); err != nil {
		oplog.Error(ctx, "failed to ping database", err)
		return err
	}

	DB = &DBContext{
		SqlDB:   database,
		Queries: sql_models.New(database),
	}

	oplog.Info(ctx, "database connected successfully")
	return nil
}

func getPostgresDsn() string {
	db := config.DB
	parts := []string{}

	if db.Host != "" {
		parts = append(parts, "host="+db.Host)
	}
	if db.User != "" {
		parts = append(parts, "user="+db.User)
	}
	if db.Password != "" {
		parts = append(parts, "password="+db.Password)
	}
	if db.Name != "" {
		parts = append(parts, "dbname="+db.Name)
	}
	if db.Port != "" {
		parts = append(parts, "port="+db.Port)
	}
	if db.SslMode != "" {
		parts = append(parts, "sslmode="+db.SslMode)
	}

	if db.SearchPath != "" {
		parts = append(parts, "search_path="+db.SearchPath)
	}

	return strings.Join(parts, " ")
}
