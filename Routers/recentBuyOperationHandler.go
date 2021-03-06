package Routers

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
	business "github.com/msrexe/portfolio-tracker/Business"
	. "github.com/msrexe/portfolio-tracker/DataAccess"
	"github.com/msrexe/portfolio-tracker/Security"
)

func GetAllOperations(c *fiber.Ctx) error {
	claims := Security.GetUserClaims(c)
	result, err := business.GetAllBuyOperations(int(claims["uid"].(float64)))
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.JSON(result)
}

func GetOperationById(c *fiber.Ctx) error {
	claims := Security.GetUserClaims(c)
	id, _ := strconv.Atoi(c.Params("id"))
	result, err := business.GetBuyOperationById(int(claims["uid"].(float64)), id)
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.JSON(result)
}

func AddOperation(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var operation RecentBuyOperation
	err := json.Unmarshal(c.Body(), &operation)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	claims := Security.GetUserClaims(c)
	if operation.UserId != int(claims["uid"].(float64)) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	result := business.AddBuyOperation(operation)
	return c.JSON(result)
}
func DeleteOperation(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var operation RecentBuyOperation
	err := json.Unmarshal(c.Body(), &operation)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	claims := Security.GetUserClaims(c)
	if operation.UserId != int(claims["uid"].(float64)) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	result := business.DeleteBuyOperation(operation)
	return c.JSON(result)
}

func UpdateOperation(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var operation RecentBuyOperation
	err := json.Unmarshal(c.Body(), &operation)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	claims := Security.GetUserClaims(c)
	if operation.UserId != int(claims["uid"].(float64)) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	result := business.UpdateBuyOperation(operation)
	return c.JSON(result)
}
