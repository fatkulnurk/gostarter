package delivery

import (
	"magicauth/internal/magiclink/domain"

	"github.com/gofiber/fiber/v2"
)

type HttpDelivery struct {
	usecase domain.IUsecase
}

func NewDeliveryHttp(usecase domain.IUsecase) *HttpDelivery {
	return &HttpDelivery{usecase: usecase}
}

func (d *HttpDelivery) HandleCreateMagicLink(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, World!",
		"status":  "success",
	})
}

func (d *HttpDelivery) HandleVerifyMagicLink(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, World!",
		"status":  "success",
	})
}
