package middlewares

import "net/http"

type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode    int
	headerWritten bool
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, 0, false}
}

func (l *LoggingResponseWriter) WriteHeader(code int) {
	if l.headerWritten {
		return
	}
	l.statusCode = code
	l.headerWritten = true
	l.ResponseWriter.WriteHeader(code)
}

func (l *LoggingResponseWriter) Write(b []byte) (int, error) {
	if !l.headerWritten {
		l.WriteHeader(http.StatusOK)
	}
	return l.ResponseWriter.Write(b)
}

func (l *LoggingResponseWriter) StatusCode() int {
	if l.statusCode == 0 {
		return http.StatusOK
	}
	return l.statusCode
}
