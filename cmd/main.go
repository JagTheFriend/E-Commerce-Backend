package main

import (
	"log/slog"
	"os"
)

func main() {
	config := config{
		addr: ":3000",
		db:   dbConfig{},
	}

	api := application{
		config: config,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := api.Start(api.mount()); err != nil {
		slog.Error("Error starting server", "error", err)
		os.Exit(1)
	}
}
