package Routers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	business "github.com/msrexe/portfolio-tracker/Business"
	. "github.com/msrexe/portfolio-tracker/DataAccess"
	"github.com/msrexe/portfolio-tracker/Security"
)

func GetAllWalletOperation(c *fiber.Ctx) error {
	claims := Security.GetUserClaims(c)
	result, err := business.GetAllWalletOperation(int(claims["uid"].(float64)))
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.JSON(result)
}

func AddWalletOperation(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var walletOperation WalletOperation
	err := json.Unmarshal(c.Body(), &walletOperation)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	result := business.AddWalletOperation(walletOperation)
	return c.JSON(result)
}

func DeleteWalletOperation(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var walletOperation WalletOperation
	err := json.Unmarshal(c.Body(), &walletOperation)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	result := business.DeleteWalletOperation(walletOperation)
	return c.JSON(result)
}
