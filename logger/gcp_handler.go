package logger

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"
)

type GCPHandler struct {
	handler             slog.Handler
	includeSourceDetail bool
}

func NewGCPLHandler(handler slog.Handler) *GCPHandler {
	return &GCPHandler{handler: handler}
}

func (h *GCPHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *GCPHandler) Handle(ctx context.Context, record slog.Record) error {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		spanCtx := span.SpanContext()

		record.AddAttrs(
			slog.String("logging.googleapis.com/trace", spanCtx.TraceID().String()),
			slog.String("logging.googleapis.com/spanId", spanCtx.SpanID().String()),
			slog.Bool("logging.googleapis.com/trace_sampled", spanCtx.IsSampled()),
		)
	}

	return h.handler.Handle(ctx, record)
}

func (h *GCPHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewGCPLHandler(h.handler.WithAttrs(attrs))
}

func (h *GCPHandler) WithGroup(name string) slog.Handler {
	return NewGCPLHandler(h.handler.WithGroup(name))
}
