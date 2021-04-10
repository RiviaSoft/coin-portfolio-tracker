package Routers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	business "github.com/msrexe/portfolio-tracker/Business"
	. "github.com/msrexe/portfolio-tracker/Core/Security"
	. "github.com/msrexe/portfolio-tracker/DataAccess"
)

func GetAllOperations(c *fiber.Ctx) error {
	return c.SendString("get all operation success")
}
func AddOperation(c *fiber.Ctx) error {
	c.Accepts("application/json")
	claims := GetUserClaims(c)
	var operation RecentBuyOperation
	err := json.Unmarshal(c.Body(), &operation)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if operation.UserId != claims["uid"] {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	result := business.AddBuyOperation(operation)
	return c.JSON(result)
}
func DeleteOperation(c *fiber.Ctx) error {
	c.Accepts("application/json")
	claims := GetUserClaims(c)
	var operation RecentBuyOperation
	err := json.Unmarshal(c.Body(), &operation)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if operation.UserId != claims["uid"] {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	result := business.DeleteBuyOperation(operation)
	return c.JSON(result)
}

func UpdateOperation(c *fiber.Ctx) error {
	c.Accepts("application/json")
	claims := GetUserClaims(c)
	var operation RecentBuyOperation
	err := json.Unmarshal(c.Body(), &operation)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if operation.UserId != claims["uid"] {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	result := business.UpdateBuyOperation(operation)
	return c.JSON(result)
}
