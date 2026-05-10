package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	ConnectDB()

	donasiRepo := NewDonasiRepository(DB)
	donasiController := NewDonasiController(donasiRepo)

	app := fiber.New()

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*", // ganti dengan domain frontend jika perlu
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: false,
		MaxAge:           43200, // 12 jam
	}))

	SetupDonasiRoutes(app, donasiController)
	app.Listen(":8080")
}
