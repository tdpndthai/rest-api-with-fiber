package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"rest-api-with-fiber/handler"
	"rest-api-with-fiber/middleware"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// Products
	product := api.Group("/product")
	product.Get("/", handler.GetAllProducts)
	product.Get("/:id", handler.GetProduct)
	product.Post("/", middleware.Protected(), handler.CreateProduct)
	product.Delete("/:id", middleware.Protected(), handler.DeleteProduct)
}
