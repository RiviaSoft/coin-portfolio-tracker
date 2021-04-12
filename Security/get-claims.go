package Security

import (
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func GetUserClaims(c *fiber.Ctx) jwt.MapClaims {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims
}
