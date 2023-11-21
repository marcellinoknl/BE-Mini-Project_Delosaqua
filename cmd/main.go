package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/marcellinoknl/-BE-Mini-Projects---Delosaqua/cmd/routes"
    "github.com/marcellinoknl/-BE-Mini-Projects---Delosaqua/database"
    "github.com/marcellinoknl/-BE-Mini-Projects---Delosaqua/handlers"
)

func main() {
    app := fiber.New()

    database.ConnectDb()

    statsMap := make(map[string]*handlers.EndpointStats)
    app.Use(handlers.StatsMiddleware(statsMap))

    routes.SetupRoutes(app, statsMap)

    app.Get("/stats", func(c *fiber.Ctx) error {
        return c.JSON(statsMap)
    })

    err := app.Listen(":3000")
    if err != nil {
        panic(err)
    }
}
