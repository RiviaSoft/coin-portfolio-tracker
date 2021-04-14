package Routers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	business "github.com/msrexe/portfolio-tracker/Business"
	. "github.com/msrexe/portfolio-tracker/DataAccess"
	Security "github.com/msrexe/portfolio-tracker/Security"
)

func GetAllArchivedOperations(c *fiber.Ctx) error {
	claims := Security.GetUserClaims(c)
	result, err := business.GetAllArchivedOperations(int(claims["uid"].(float64)))
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.JSON(result)
}

func AddArchivedOperation(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var operation ArchivedOperation
	err := json.Unmarshal(c.Body(), &operation)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	claims := Security.GetUserClaims(c)
	if operation.UserId != int(claims["uid"].(float64)) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	result := business.AddArchivedOperation(operation)
	return c.JSON(result)
}

func DeleteArchivedOperation(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var operation ArchivedOperation
	err := json.Unmarshal(c.Body(), &operation)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	claims := Security.GetUserClaims(c)
	if operation.UserId != int(claims["uid"].(float64)) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	result := business.DeleteArchivedOperation(operation)
	return c.JSON(result)
}
