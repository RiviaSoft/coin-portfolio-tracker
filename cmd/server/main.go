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
	app := fiber.New()
	Routers.SetupRoutes(app)
	Security.SetupMW(app)
	app.Listen(port)
}
