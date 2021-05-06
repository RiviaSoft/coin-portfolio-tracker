package Routers

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
	business "github.com/msrexe/portfolio-tracker/Business"
	. "github.com/msrexe/portfolio-tracker/DataAccess"
)

func GetAllWalletOperation(c *fiber.Ctx) error {
	//claims := Security.GetUserClaims(c)
	id, _ := strconv.Atoi(c.Query("id"))
	result, err := business.GetAllWalletOperation(id)
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
