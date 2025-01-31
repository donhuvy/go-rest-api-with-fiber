package main

import (
	"github.com/donhuvy/go-rest-api-with-fiber/database"
	"github.com/donhuvy/go-rest-api-with-fiber/router"
	"github.com/gofiber/fiber" // import the fiber package
	"github.com/gofiber/fiber/middleware"
	_ "github.com/lib/pq"
	"log"
)

func main() { // entry point to our program
	// Connect to database.
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New() // call the New() method - used to instantiate a new Fiber App
	app.Use(middleware.Logger())
	router.SetupRoutes(app)
	app.Listen(3000) // Listen/Serve the new Fiber app on port 3000.
}
