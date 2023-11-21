package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/marcellinoknl/-BE-Mini-Projects---Delosaqua/cmd/routes"
    "github.com/marcellinoknl/-BE-Mini-Projects---Delosaqua/database"
    "github.com/marcellinoknl/-BE-Mini-Projects---Delosaqua/handlers"
)

func main() {
    // Create a new Fiber web application.
    app := fiber.New()

    // Connect to the PostgreSQL database using the database package.
    database.ConnectDb()

    // Create a map to store endpoint statistics using handlers package.
    statsMap := make(map[string]*handlers.EndpointStats)

    // Use a middleware to record statistics for incoming HTTP requests.
    app.Use(handlers.StatsMiddleware(statsMap))

    // Set up the application routes using the routes package.
    routes.SetupRoutes(app, statsMap)

    // Define a route to expose statistics data as JSON.
    app.Get("/stats", func(c *fiber.Ctx) error {
        return c.JSON(statsMap)
    })

    // Start the Fiber web server on port 3000.
    err := app.Listen(":3000")
    if err != nil {
        panic(err)
    }
}
