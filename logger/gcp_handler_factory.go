package logger

import (
	"log/slog"
	"os"
)

func NewGCPHandler(level, projectId string) slog.Handler {
	return NewGCPLHandler(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			ReplaceAttr: ReplaceAttr,
			AddSource:   true,
			Level:       ConvertToSlogLevel(level),
		}),
		projectId,
	)
}

func ReplaceAttr(groups []string, attr slog.Attr) slog.Attr {
	noGroups := len(groups) == 0

	switch {
	case noGroups && attr.Key == slog.TimeKey:
		return attr
	case noGroups && attr.Key == slog.LevelKey:
		logLevel, ok := attr.Value.Any().(slog.Level)
		if !ok {
			return attr
		}
		return slog.String(attrSeverity, getSeverity(logLevel))
	case noGroups && attr.Key == slog.SourceKey:
		source, ok := attr.Value.Any().(*slog.Source)
		if !ok || source == nil {
			return attr
		}
		return slog.Any(sourceLocationKey, source)
	case noGroups && attr.Key == slog.MessageKey:
		return slog.String(attrMessage, attr.Value.String())
	default:
		return attr
	}
}

func getSeverity(level slog.Level) string {
	switch level {
	case slog.LevelDebug:
		return levelDebug
	case slog.LevelInfo:
		return levelInfo
	case slog.LevelWarn:
		return levelWarn
	case slog.LevelError:
		return levelError
	default:
		return levelDefault
	}
}

func ConvertToSlogLevel(level string) slog.Level {
	switch level {
	case levelDebug:
		return slog.LevelDebug
	case levelInfo:
		return slog.LevelInfo
	case levelWarn:
		return slog.LevelWarn
	case levelError:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
