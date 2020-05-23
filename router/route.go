package router

import (
	"github.com/donhuvy/go-rest-api-with-fiber/handler"
	"github.com/donhuvy/go-rest-api-with-fiber/middleware"
	"github.com/gofiber/fiber"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// Middleware.
	api := app.Group("/api", middleware.AuthReq())
	// Routes.
	api.Get("/", handler.GetAllProducts)
	api.Get("/:id", handler.GetSingleProduct)
	api.Post("/", handler.CreateProduct)
	api.Delete("/:id", handler.DeleteProduct)
}
