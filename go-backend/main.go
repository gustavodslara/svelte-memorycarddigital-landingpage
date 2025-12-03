package main

import (
	"go-backend/database"
	"go-backend/handlers"
	"go-backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.User{}, &models.GameSave{})

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:5173,https://memorycard.digital", // Adjust as needed
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
	}))

	app.Post("/api/register", handlers.Register)
	app.Post("/api/login", handlers.Login)
	app.Get("/api/user", handlers.Me)

	app.Post("/api/saves", handlers.UploadSave)
	app.Get("/api/saves", handlers.GetSaves)
	app.Get("/api/saves/:id/download", handlers.DownloadSave)

	app.Get("/api/subscription", handlers.GetSubscription)
	app.Post("/api/subscription", handlers.Subscribe)

	app.Listen(":8000")
}
