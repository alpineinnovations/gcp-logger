package middlewares

import (
	"fmt"
	log "github.com/alpineinnovations/gcp-logger/logger"
	"log/slog"
	"net/http"
	"time"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.FromCtx(r.Context())

		start := time.Now()
		logRequestReceived(logger, r)

		writer := NewLoggingResponseWriter(w)
		next.ServeHTTP(writer, r)

		elapsedTimeInSec := time.Since(start).Seconds()
		logRequestServed(logger, writer, r, elapsedTimeInSec)
	})
}

func logRequestReceived(logger *slog.Logger, r *http.Request) {
	logger.Info("Request",
		slog.Group(log.KeyHttpRequest,
			slog.String(log.KeyRequestMethod, r.Method),
			slog.String(log.KeyRequestUrl, r.URL.String()),
			slog.String(log.KeyRemoteIp, r.RemoteAddr),
			slog.String(log.KeyUserAgent, r.UserAgent()),
			slog.String(log.KeyReferer, r.Referer()),
			slog.String(log.KeyProtocol, r.Proto),
		),
	)
}

func logRequestServed(logger *slog.Logger, w *LoggingResponseWriter, r *http.Request, latencyInSec float64) {

	logger.Info("Request.Served",
		slog.Group(log.KeyHttpRequest,
			slog.String(log.KeyRequestMethod, r.Method),
			slog.String(log.KeyRequestUrl, r.URL.String()),
			slog.String(log.KeyRemoteIp, r.RemoteAddr),
			slog.String(log.KeyUserAgent, r.UserAgent()),
			slog.String(log.KeyReferer, r.Referer()),
			slog.String(log.KeyProtocol, r.Proto),
			slog.String(log.KeyLatency, fmt.Sprintf("%fs", latencyInSec)),
			slog.Int(log.KeyStatus, w.StatusCode()),
		),
	)
}
