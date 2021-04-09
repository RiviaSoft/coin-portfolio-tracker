package Routers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	business "github.com/msrexe/portfolio-tracker/Business"
	. "github.com/msrexe/portfolio-tracker/DataAccess"
	. "github.com/msrexe/portfolio-tracker/Middlewares/Security"
)

func GetAllArchivedOperations(c *fiber.Ctx) error {
	c.Accepts("application/json")
	claims := GetUserClaims(c)
	var operation ArchivedOperation
	err := json.Unmarshal(c.Body(), &operation)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	result := business.GetAllArchivedOperations(int(claims["uid"].(float64)))
	return c.JSON(result)
}
func AddArchivedOperation(c *fiber.Ctx) error {
	c.Accepts("application/json")
	claims := GetUserClaims(c)
	var operation ArchivedOperation
	err := json.Unmarshal(c.Body(), &operation)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	result := business.AddArchivedOperation(int(claims["uid"].(float64)), operation)
	return c.JSON(result)
}
func DeleteArchivedOperation(c *fiber.Ctx) error {
	c.Accepts("application/json")
	claims := GetUserClaims(c)
	var operation ArchivedOperation
	err := json.Unmarshal(c.Body(), &operation)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	result := business.DeleteArchivedOperation(int(claims["uid"].(float64)), operation)
	return c.JSON(result)
}
