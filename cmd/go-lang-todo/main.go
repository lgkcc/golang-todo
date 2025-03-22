package main

import (
	"go-lang-todo/internal/config"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

const envLocal = "local"
const envDev = "dev"
const envProd = "prod"

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("failed to load .env")
	}

	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting app", slog.String("env", cfg.Env), slog.String("version", "123"))
	log.Debug("debug messages are enabled")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
