package Routers

import "github.com/gofiber/fiber/v2"

func GetAllArchivedOperations(c *fiber.Ctx) error {
	return c.SendString("Test başarılı :D")
}
func AddArchivedOperation(c *fiber.Ctx) error {
	return c.SendString("Test başarılı :D")
}
func DeleteArchivedOperation(c *fiber.Ctx) error {
	return c.SendString("Test başarılı :D")
}
func UpdateArchivedOperation(c *fiber.Ctx) error {
	return c.SendString("Test başarılı :D")
}
