package logger

import (
	"context"
	"fmt"
	"log/slog"

	"go.opentelemetry.io/otel/trace"
)

type GCPHandler struct {
	handler   slog.Handler
	projectId string
}

func NewGCPLHandler(handler slog.Handler, projectId string) *GCPHandler {
	return &GCPHandler{handler: handler, projectId: projectId}
}

func (h *GCPHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *GCPHandler) Handle(ctx context.Context, record slog.Record) error {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		spanCtx := span.SpanContext()

		traceId := fmt.Sprintf("projects/%s/traces/%s", h.projectId, spanCtx.TraceID().String())
		spanId := spanCtx.SpanID().String()

		record.AddAttrs(
			slog.String("logging.googleapis.com/trace", traceId),
			slog.String("logging.googleapis.com/spanId", spanId),
			slog.Bool("logging.googleapis.com/trace_sampled", spanCtx.IsSampled()),
		)
	}

	return h.handler.Handle(ctx, record)
}

func (h *GCPHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewGCPLHandler(h.handler.WithAttrs(attrs), h.projectId)
}

func (h *GCPHandler) WithGroup(name string) slog.Handler {
	return NewGCPLHandler(h.handler.WithGroup(name), h.projectId)
}
