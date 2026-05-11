package logging

import (
	"log/slog"
	"os"
	"strings"
)

func NewLogger() *slog.Logger {
	logLevel := slog.LevelInfo

	if strings.ToLower(os.Getenv("LOG_LEVEL")) == "debug" {
		logLevel = slog.LevelDebug
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)
	logger := slog.New(handler)

	if logLevel == slog.LevelDebug {
		logger.Info("Debug logging enabled")
	} else {
		logger.Info("Info logging enabled")
	}

	return logger
}
