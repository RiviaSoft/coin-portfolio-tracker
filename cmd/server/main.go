package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/msrexe/portfolio-tracker/Middlewares/Security"
	"github.com/msrexe/portfolio-tracker/Routers"
)

const (
	port = ":8000"
)

func main() {

}

func init() {
	app := fiber.New()
	api := app.Group("/api")

	Security.SetupMW(api)
	Routers.SetupRoutes(api)

	app.Listen(port)
}
