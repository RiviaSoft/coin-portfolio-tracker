package Routers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	business "github.com/msrexe/portfolio-tracker/Business"
	. "github.com/msrexe/portfolio-tracker/DataAccess"
	. "github.com/msrexe/portfolio-tracker/Middlewares/Security"
)

func GetAllOperations(c *fiber.Ctx) error {
	c.Accepts("application/json")
	return c.SendString("get all operation success")
}
func AddOperation(c *fiber.Ctx) error {
	c.Accepts("application/json")
	claims := GetUser(c)
	var operation RecentBuyOperation
	err := json.Unmarshal(c.Body(), &operation)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	result := business.AddBuyOperation(int(claims["uid"].(float64)), operation)
	return c.JSON(result)
}
func DeleteOperation(c *fiber.Ctx) error {
	c.Accepts("application/json")
	claims := GetUser(c)
	var operation RecentBuyOperation
	err := json.Unmarshal(c.Body(), &operation)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	result := business.DeleteBuyOperation(int(claims["uid"].(float64)), operation)
	return c.JSON(result)
}

func UpdateOperation(c *fiber.Ctx) error {
	c.Accepts("application/json")
	claims := GetUser(c)
	var operation RecentBuyOperation
	err := json.Unmarshal(c.Body(), &operation)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	result := business.UpdateBuyOperation(int(claims["uid"].(float64)), operation)
	return c.JSON(result)
}
