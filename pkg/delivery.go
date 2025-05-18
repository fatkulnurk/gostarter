package pkg

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
)

type Delivery struct {
	HTTP   *fiber.App
	Worker *asynq.ServeMux
}
