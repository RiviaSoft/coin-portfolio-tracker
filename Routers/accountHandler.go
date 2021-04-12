package Routers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/msrexe/portfolio-tracker/Security"
)

func AccountHandler(c *fiber.Ctx) error {
	claims := Security.GetUserClaims(c)
	expTime := fmt.Sprintf("%f", claims["exp"])
	msg := "Active account :  " + claims["uname"].(string) + "\nMail adresi : " + claims["email"].(string) + "\n Token ömrü : " + expTime
	return c.SendString(msg)
}
