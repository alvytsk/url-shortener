package main

import (
	"alvytsk/url-shortener/internal/config"
	"alvytsk/url-shortener/internal/storage/sqlite"
	"log/slog"
	"os"
)

func main() {
	cfg := config.LoadConfig()
	
	logger := setupLogger(cfg)
	logger.Info("init server", "env", cfg.Env)
	logger.Debug("debug message")

	storage, err := sqlite.NewStorage(cfg.StoragePath)
	if err != nil {
		logger.Error("failed to init storage", "error", err)
		os.Exit(1)
	}

	_ = storage


	//TODO: init router -> chi, render (chi render)

	//TODO: run server
}

func setupLogger(cfg *config.Config) *slog.Logger {
	var logger *slog.Logger

	if(cfg.Env == "local") {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}