package main

import (
	"log/slog"

	"github.com/damirchikk/TestProj2/internal/pkg/logger"
)

func main() {
	slogger := logger.New()
	slog.SetDefault(slogger)
	slog.SetDefault(slogger)
	slog.SetDefault(slogger)
}
