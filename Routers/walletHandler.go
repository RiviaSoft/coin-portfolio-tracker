package Routers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	business "github.com/msrexe/portfolio-tracker/Business"
	. "github.com/msrexe/portfolio-tracker/DataAccess"
	"github.com/msrexe/portfolio-tracker/Security"
)

func GetAllWallets(c *fiber.Ctx) error {
	claims := Security.GetUserClaims(c)
	result, err := business.GetWalletByUserId(int(claims["uid"].(float64)))
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.JSON(result)
}

func GetWalletById(c *fiber.Ctx) error {
	claims := Security.GetUserClaims(c)
	result, err := business.GetWalletByName(int(claims["uid"].(float64)), c.Params("wallet"))
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.JSON(result)
}

func AddWallet(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var wallet Wallet
	err := json.Unmarshal(c.Body(), &wallet)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	claims := Security.GetUserClaims(c)
	if wallet.UserId != int(claims["uid"].(float64)) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	result := business.AddWallet(wallet)
	return c.JSON(result)
}

func DeleteWallet(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var wallet Wallet
	err := json.Unmarshal(c.Body(), &wallet)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	claims := Security.GetUserClaims(c)
	if wallet.UserId != int(claims["uid"].(float64)) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	result := business.DeleteWallet(wallet)
	return c.JSON(result)
}

func DeleteAllWallets(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var wallet Wallet
	err := json.Unmarshal(c.Body(), &wallet)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	claims := Security.GetUserClaims(c)
	if wallet.UserId != int(claims["uid"].(float64)) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	result := business.DeleteAllWallet(wallet.UserId, wallet.Name)
	return c.JSON(result)
}
