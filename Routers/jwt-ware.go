package Routers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	business "github.com/msrexe/portfolio-tracker/Business"
	. "github.com/msrexe/portfolio-tracker/Core/EnvVariables"
	. "github.com/msrexe/portfolio-tracker/DataAccess"
	"github.com/msrexe/portfolio-tracker/Security"
)

func SetupMW(api fiber.Router) {
	auth := api.Group("/auth")
	auth.Post("/login", loginHandler)
	auth.Post("/register", registerHandler)

	//JWTware
	api.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(GoDotEnvVariable("SIGNING_KEY")),
	}))

}

func registerHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var user UserModel
	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	result := business.AddUser(user)
	return c.JSON(result)
}

func loginHandler(c *fiber.Ctx) error {
	email := c.FormValue("email")
	pass := c.FormValue("password")

	user, err := business.GetUser(email)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	// Throws Unauthorized error
	if !Security.CheckPasswordHash(pass, user.PasswordHash) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	t, err := Security.CreateToken(user.Id, user.Name, user.Email)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"token": t})
}
