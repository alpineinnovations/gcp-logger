package middlewares

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	log "github.com/alpineinnovations/gcp-logger/logger"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.FromCtx(r.Context())

		start := time.Now()
		logRequestReceived(logger, r)

		next.ServeHTTP(w, r)

		elapsedTimeInSec := time.Since(start).Seconds()
		logRequestServed(logger, w, r, elapsedTimeInSec)
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

func logRequestServed(logger *slog.Logger, w http.ResponseWriter, r *http.Request, latencyInSec float64) {

	logger.Info("Request.Served",
		slog.Group(log.KeyHttpRequest,
			slog.String(log.KeyRequestMethod, r.Method),
			slog.String(log.KeyRequestUrl, r.URL.String()),
			slog.String(log.KeyRemoteIp, r.RemoteAddr),
			slog.String(log.KeyUserAgent, r.UserAgent()),
			slog.String(log.KeyReferer, r.Referer()),
			slog.String(log.KeyProtocol, r.Proto),
			slog.String(log.KeyLatency, fmt.Sprintf("%fs", latencyInSec)),
		),
	)
}
