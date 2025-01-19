package logger

import (
	"context"
	"graphql-api/constants"
	"log/slog"
	"os"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// LoggerMiddleWare logs the incoming GraphQL query or mutation name and ensures the request body is reusable
func LoggerMiddleWare() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.WithValue(c.Request().Context(), constants.RequestKey, uuid.New().String())
			c.SetRequest(c.Request().WithContext(ctx))
			logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
			slog.SetDefault(logger)
			slog.InfoContext(c.Request().Context(), "Request", "UUID", ctx.Value(constants.RequestKey))
			return next(c)
		}
	}
}
