package logger

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func LogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Generate unique requestID
		requestID := uuid.New().String()

		// Create a pre-warmed logger with requestID
		reqLogger := WithRequestID(requestID)

		// Inject into context
		ctx := InjectLogger(c.Request().Context(), reqLogger)
		ctx = context.WithValue(ctx, "requestID", requestID)
		c.SetRequest(c.Request().WithContext(ctx))

		// Log receipt
		Log(ctx, zerolog.InfoLevel, "Received request", map[string]any{
			"method": c.Request().Method,
			"url":    c.Request().URL.String(),
		})

		return next(c)
	}
}
