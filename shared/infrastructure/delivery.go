package infrastructure

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
)

// Delivery manages all input/output mechanisms for the application
// It follows the clean architecture pattern as the driver/delivery layer
// This includes HTTP servers, task handlers, and scheduled jobs
type Delivery struct {
	HTTP     *fiber.App       // HTTP server for handling web requests
	Task     *asynq.ServeMux  // Task handler for processing background jobs
	Schedule *asynq.Scheduler // Scheduler for managing periodic tasks
}
