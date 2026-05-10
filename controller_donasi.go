package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type DonasiController struct {
	DonasiRepo *DonasiRepository
}

func NewDonasiController(repo *DonasiRepository) *DonasiController {
	return &DonasiController{DonasiRepo: repo}
}

func (c *DonasiController) GetAllDonations(ctx *fiber.Ctx) error {
	donations, err := c.DonasiRepo.GetAllDonations()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(donations)
}

func (c *DonasiController) GetDonationByID(ctx *fiber.Ctx) error {
	id_donasi, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || id_donasi <= 0 {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	donation, err := c.DonasiRepo.GetDonationByID(id_donasi)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Donation Not Found",
		})
	}

	return ctx.Status(http.StatusOK).JSON(donation)
}

func (c *DonasiController) CreateDonation(ctx *fiber.Ctx) error {
	var donation Donation

	if err := ctx.BodyParser(&donation); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// VALIDATION
	if strings.TrimSpace(donation.NamaDonatur) == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Nama donatur wajib diisi",
		})
	}

	if donation.Nominal <= 0 {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Nominal harus lebih dari 0",
		})
	}

	if donation.TanggalDonasi.IsZero() {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Tanggal donasi wajib diisi",
		})
	}

	if strings.TrimSpace(donation.MetodePembayaran) == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Metode pembayaran wajib diisi",
		})
	}

	if strings.TrimSpace(donation.Status) == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Status wajib diisi",
		})
	}

	if err := c.DonasiRepo.CreateDonation(&donation); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(donation)
}

func (c *DonasiController) UpdateDonation(ctx *fiber.Ctx) error {
	id_donasi, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || id_donasi <= 0 {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var donation Donation

	if err := ctx.BodyParser(&donation); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// VALIDATION
	if strings.TrimSpace(donation.NamaDonatur) == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Nama donatur wajib diisi",
		})
	}

	if donation.Nominal <= 0 {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Nominal harus lebih dari 0",
		})
	}

	if donation.TanggalDonasi.IsZero() {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Tanggal donasi wajib diisi",
		})
	}

	if strings.TrimSpace(donation.MetodePembayaran) == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Metode pembayaran wajib diisi",
		})
	}

	if strings.TrimSpace(donation.Status) == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Status wajib diisi",
		})
	}

	donation.IDDonasi = id_donasi

	if err := c.DonasiRepo.UpdateDonation(&donation); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(donation)
}

func (c *DonasiController) DeleteDonation(ctx *fiber.Ctx) error {
	id_donasi, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || id_donasi <= 0 {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	if err := c.DonasiRepo.DeleteDonation(id_donasi); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Donation deleted successfully",
	})
}