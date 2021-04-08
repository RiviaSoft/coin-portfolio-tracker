package Routers

import (
	"github.com/gofiber/fiber/v2"
)

func TestHandler(c *fiber.Ctx) error {
	return c.SendString("Test başarılı :D")
}
