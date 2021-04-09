package Routers

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(api fiber.Router) {
	api.Get("/test", TestHandler)

	api.Get("/account", AccountHandler)
	//Diğer account operasyonları jwt-ware içinde

	//Recent Buy Operations
	//api.Get("/operations/get/id::id", GetByOperationId)
	api.Get("/operations/getall", GetAllOperations)
	api.Post("/operations/add", AddOperation)
	api.Post("/operations/delete", DeleteOperation)
	api.Post("/operations/update", UpdateOperation)

	//Archived Buy Operations
	api.Get("/archivedOperations/getall", GetAllArchivedOperations)
	//api.Get("/archivedOperations/get/id::id", GetByOperationId)
	api.Post("/archivedOperations/add", AddArchivedOperation)
	api.Post("/archivedOperations/delete", DeleteArchivedOperation)
}
