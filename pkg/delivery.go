package pkg

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
)

type Delivery struct {
	HTTP     *fiber.App
	Task     *asynq.ServeMux
	Schedule *asynq.Scheduler
}
