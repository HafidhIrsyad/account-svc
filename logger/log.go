package logger

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Logger is the base zerolog.Logger
var Logger zerolog.Logger

// key type to avoid collisions
type ctxKeyLogger struct{}

func init() {
	// Use ConsoleWriter for pretty terminal output, including timestamp
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	Logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
}

// WithRequestID returns a logger with requestID field
func WithRequestID(requestID string) zerolog.Logger {
	return Logger.With().Str("requestID", requestID).Logger()
}

// InjectLogger stores the given logger into ctx
func InjectLogger(ctx context.Context, l zerolog.Logger) context.Context {
	return context.WithValue(ctx, ctxKeyLogger{}, l)
}

// CtxLogger retrieves the logger from ctx or returns base Logger
func CtxLogger(ctx context.Context) zerolog.Logger {
	if l, ok := ctx.Value(ctxKeyLogger{}).(zerolog.Logger); ok {
		return l
	}
	return Logger
}

// Log writes a message at given level with dynamic fields
func Log(ctx context.Context, level zerolog.Level, msg string, fields map[string]any) {
	l := CtxLogger(ctx)
	event := l.WithLevel(level)
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(msg)
}
