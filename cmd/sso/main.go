package main

import (
	"log/slog"
	"os"
	"sso/internal/config"
	"sso/internal/lib/logger/handlers/slogpretty"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TODO: инициализация объекта конфига
	cfg := config.MustLoad()

	// TODO: инициализация объекта логгера
	log := setupLogger(cfg.Env)
	log.Info("start application", slog.Any("config", cfg))
	// sl.Err(err)

	// TODO: инициализация объекта приложения (app)

	// TODO: запуск приложения gRPC-сервер приложения (app.Run())
}

func setupLogger(env string) *slog.Logger {
	// TODO: настройка логгера
	var logger *slog.Logger
	switch env {
	case envLocal:
		logger = setupPrettyLogger()
	case envDev:
		logger = slog.New(
			slog.NewJSONHandler(
				os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug},
			),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(
				os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo},
			),
		)
	}
	return logger
}

func setupPrettyLogger() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}
	handler := opts.NewPrettyHandler(os.Stdout)
	return slog.New(handler)
}
