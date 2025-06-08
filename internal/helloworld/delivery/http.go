package delivery

import (
	"github.com/fatkulnurk/gostarter/internal/helloworld/domain"

	"github.com/gofiber/fiber/v2"
)

type HttpDelivery struct {
	service domain.Service
}

func NewDeliveryHttp(service domain.Service) *HttpDelivery {
	return &HttpDelivery{service: service}
}

func (d *HttpDelivery) HandleHelloWorld(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, World!",
		"status":  "success",
	})
}

func (d *HttpDelivery) HandleExampleApi(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, World!",
		"status":  "success",
	})
}
