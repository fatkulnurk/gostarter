package middleware

import (
	"context"
	"github.com/fatkulnurk/gostarter/pkg/logging"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Record start time
		start := time.Now()

		// Defer logging until after request is processed
		defer func() {
			// Calculate request duration
			duration := time.Since(start)

			// Collect request details
			statusCode := c.Response().StatusCode()
			method := c.Method()
			path := c.Path()
			ip := c.IP()
			userAgent := string(c.Request().Header.UserAgent())

			// Check if there's a mismatch between actual status and what's being logged
			if c.Response().StatusCode() != statusCode {
				logging.Warning(context.Background(), "Status code mismatch", logging.NewField("actual_status", c.Response().StatusCode()))
			}

			// Log request information
			logging.Info(context.Background(), "Incoming request",
				logging.NewField("method", method),
				logging.NewField("path", path),
				logging.NewField("status", statusCode),
				logging.NewField("ip", ip),
				logging.NewField("user_agent", userAgent),
				logging.NewField("latency", duration),
			)
		}()

		// Process request
		return c.Next()
	}
}
