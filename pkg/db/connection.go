package db

import (
	"context"
	"errors"
	"log/slog"
	"oxid-gateway-admin-api/pkg/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func Connect() {
	config, err := pgxpool.ParseConfig(config.GetEnvOrPanic(config.DatabaseURL))

	if err != nil {
		slog.Error("Invalid database URL", "eror", err)
        panic(err)
	}

	pool, err = pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		slog.Error("Failed to connect to database pool", "eror", err)
        panic(err)
	}

	slog.Info("Connected to database")
}

func GetConection() (*pgxpool.Conn, error) {
	if pool == nil {
        return nil, errors.New("Not connected to database")
	}

	conn, err := pool.Acquire(context.Background())

	if err != nil {
		slog.Error("Failed to aquire database connection", "eror", err)
        return nil, err
	}

	return conn, nil
}

func Close() {
	if pool != nil {
		pool.Close()
	}

	pool = nil
}
