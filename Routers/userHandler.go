package Routers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	business "github.com/msrexe/portfolio-tracker/Business"
	. "github.com/msrexe/portfolio-tracker/DataAccess"
	"github.com/msrexe/portfolio-tracker/Security"
)

func GetCurrentUser(c *fiber.Ctx) error {
	claims := Security.GetUserClaims(c)
	result, err := business.GetUser(claims["email"].(string))
	if err != nil {
		return c.SendString(err.Error())
	}
	result.PasswordHash = ""
	return c.JSON(result)
}

func UpdateUser(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var user UserModel
	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	claims := Security.GetUserClaims(c)
	if user.Id != int(claims["uid"].(float64)) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	result := business.UpdateUser(user)
	return c.JSON(result)
}

func DeleteUser(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var user User
	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	claims := Security.GetUserClaims(c)
	if user.Id != int(claims["uid"].(float64)) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	result := business.DeleteUser(user)
	return c.JSON(result)
}
