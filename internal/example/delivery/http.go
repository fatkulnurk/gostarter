package delivery

import (
	"github.com/fatkulnurk/gostarter/internal/example/domain"

	"github.com/gofiber/fiber/v2"
)

type HttpDelivery struct {
	usecase domain.IUsecase
}

func NewDeliveryHttp(usecase domain.IUsecase) *HttpDelivery {
	return &HttpDelivery{usecase: usecase}
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
