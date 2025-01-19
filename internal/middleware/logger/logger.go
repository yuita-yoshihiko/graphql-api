package logger

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
)

// LoggerMiddleWare logs the incoming GraphQL query or mutation name and ensures the request body is reusable
func LoggerMiddleWare() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
			slog.SetDefault(logger)
			return next(c)
		}
	}
}
