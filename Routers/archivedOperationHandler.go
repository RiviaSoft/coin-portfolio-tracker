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
	result, _ := business.GetAllArchivedOperations(int(claims["uid"].(float64)))
	return c.JSON(result)
}

func AddArchivedOperation(c *fiber.Ctx) error {
	c.Accepts("application/json")
	claims := Security.GetUserClaims(c)
	var operation ArchivedOperation
	err := json.Unmarshal(c.Body(), &operation)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if operation.UserId != claims["uid"] {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	result := business.AddArchivedOperation(operation)
	return c.JSON(result)
}

func DeleteArchivedOperation(c *fiber.Ctx) error {
	c.Accepts("application/json")
	claims := Security.GetUserClaims(c)
	var operation ArchivedOperation
	err := json.Unmarshal(c.Body(), &operation)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if operation.UserId != claims["uid"] {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	result := business.DeleteArchivedOperation(operation)
	return c.JSON(result)
}
