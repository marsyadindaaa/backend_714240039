package main

import (
	"github.com/gofiber/fiber/v2"
)

func SetupDonasiRoutes(app *fiber.App, donasiController *DonasiController) {
	donasi := app.Group("/donations")

	donasi.Get("/", donasiController.GetAllDonations)
	donasi.Get("/:id", donasiController.GetDonationByID)
	donasi.Post("/", donasiController.CreateDonation)
	donasi.Put("/:id", donasiController.UpdateDonation)
	donasi.Delete("/:id", donasiController.DeleteDonation)
}