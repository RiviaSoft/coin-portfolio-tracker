package Routers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/msrexe/portfolio-tracker/Middlewares/Security"
)

func AccountHandler(c *fiber.Ctx) error {
	claims, err := Security.GetUser(c)
	if err != nil {
		log.Println(err.Error())
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	expTime := fmt.Sprintf("%f", claims["exp"])
	msg := "Active account :  " + claims["uname"].(string) + "\nMail adresi : " + claims["email"].(string) + "\n Token ömrü : " + expTime
	return c.SendString(msg)
}
