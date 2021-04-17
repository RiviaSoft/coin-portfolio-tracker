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
	api.Get("/archivedoperations/getall", GetAllArchivedOperations)
	//api.Get("/archivedOperations/get/id::id", GetByOperationId)
	api.Post("/archivedoperations/add", AddArchivedOperation)
	api.Post("/archivedOperations/delete", DeleteArchivedOperation)

	//Wallet
	api.Get("/wallets/getall", GetAllWallets)
	api.Get("/wallets/getbyid", GetWalletById)
	api.Post("/wallets/add", AddWallet)
	api.Post("/wallets/delete", DeleteWallet)
	//update eklenecek

	//Wallet Operations
	api.Get("/walletoperations/getall", GetAllWalletOperation)
	api.Post("/walletoperations/add", AddWalletOperation)
	api.Post("/walletoperations/delete", DeleteWallet)
}
