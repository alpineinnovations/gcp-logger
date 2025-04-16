package logger

import (
	"context"
	"log/slog"
)

func FromCtx(ctx context.Context) *slog.Logger {
	if ctx == nil {
		return slog.Default()
	}
	logger, ok := ctx.Value(AppLogger).(*slog.Logger)
	if !ok {
		return slog.Default()
	}
	return logger
}

func LoggerContext(ctx context.Context, userID string) context.Context {
	logger := slog.Default().With(KeyUser, userID)
	return context.WithValue(ctx, AppLogger, logger)
}
