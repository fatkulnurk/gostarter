package middleware

import (
	"time"

	"github.com/fatkulnurk/gostarter/pkg/logging"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
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
				logging.Warn("Status code mismatch",
					zap.Int("actual_status", c.Response().StatusCode()),
					zap.Int("logged_status", statusCode))
			}

			// Log request information
			logging.Info("Incoming request",
				zap.String("method", method),
				zap.String("path", path),
				zap.Int("status", c.Response().StatusCode()), // Get status directly from response
				zap.String("ip", ip),
				zap.String("user_agent", userAgent),
				zap.Duration("latency", duration),
			)
		}()

		// Process request
		return c.Next()
	}
}
