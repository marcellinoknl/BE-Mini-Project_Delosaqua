package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/marcellinoknl/-BE-Mini-Projects---Delosaqua/handlers"
)

func SetupRoutes(app *fiber.App, statsMap map[string]*handlers.EndpointStats) {
    v1 := app.Group("/v1")

    // Farm routes
    v1.Get("/farms-delosaqua/viewAll", handlers.ListFarms)
    v1.Post("/farms-delosaqua/manage/store", handlers.CreateFarm)
    v1.Put("/farms-delosaqua/manage/update/:id", handlers.UpdateFarm)
    v1.Delete("/farms-delosaqua/delete/:id", handlers.DeleteFarm)
    v1.Get("/farms-delosaqua/viewById/:id", handlers.GetFarmByID)

    // Pond routes
    v1.Post("/ponds-delosaqua/manage/store", handlers.CreatePond)
    v1.Put("/ponds-delosaqua/manage/update/:id", handlers.UpdatePond)
    v1.Delete("/ponds-delosaqua/delete/:id", handlers.DeletePond)
    v1.Get("/ponds-delosaqua/viewById/:id", handlers.GetPondByID)
    v1.Get("/ponds-delosaqua/viewAll", handlers.ListPonds)

    // API statistics route
    v1.Get("/get-stats", func(c *fiber.Ctx) error {
        return c.JSON(statsMap)
    })
}
