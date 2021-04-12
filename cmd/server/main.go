package main

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/msrexe/portfolio-tracker/Core/EnvVariables"
	"github.com/msrexe/portfolio-tracker/Routers"
)

func main() {

}

func init() {
	app := fiber.New()
	api := app.Group("/api")

	Routers.SetupMW(api)
	Routers.SetupRoutes(api)

	app.Listen(":" + GoDotEnvVariable("PORT"))
}
