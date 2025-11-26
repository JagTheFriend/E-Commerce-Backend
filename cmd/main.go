package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/JagTheFriend/ecommerce/internal/env"
	"github.com/jackc/pgx/v5"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := config{
		addr: ":3000",
		db: dbConfig{
			dsn: env.GetEnv("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}

	// Database connection
	ctx := context.Background()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	conn, err := pgx.Connect(ctx, config.db.dsn)
	defer func() {
		if err := conn.Close(ctx); err != nil {
			slog.Error("Failed to close database connection", "error", err)
			os.Exit(1)
		}
	}()

	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}

	slog.Info("Connected to database", "dsn", config.db.dsn)

	api := application{
		config: config,
	}

	if err := api.Start(api.mount()); err != nil {
		slog.Error("Error starting server", "error", err)
		os.Exit(1)
	}
}
