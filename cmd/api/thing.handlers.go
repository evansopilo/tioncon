package main

import (
	"log"
	"strconv"

	"github.com/evansopilo/tioncon/database"
	"github.com/evansopilo/tioncon/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (app *application) InsertThing(ctx *fiber.Ctx) error {
	var thing models.IThing = models.NewThing()
	if err := ctx.BodyParser(&thing); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	thing.SetID(uuid.New().String())
	if err := app.Things.Insert(thing); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "id": thing.GetThing().ID})
}

func (app *application) ReadThing(ctx *fiber.Ctx) error {
	thing, err := app.Things.Read(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": thing})
}

func (app *application) FetchThing(ctx *fiber.Ctx) error {
	var opts database.FetchOptions
	if ctx.Params("id") != "" {
		opts.ID = ctx.Params("id")
	}
	if ctx.Params("device_id") != "" {
		opts.DeviceID = ctx.Params("device_id")
	}

	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil {
		log.Println(err)
	}

	limit, err := strconv.Atoi(ctx.Query("limit", "15"))
	if err != nil {
		log.Println(err)
	}

	things, err := app.Things.Fetch(opts, int64(page), int64(limit))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": things})
}

func (app *application) RemoveThing(ctx *fiber.Ctx) error {
	if err := app.Things.Remove(ctx.Params("id")); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "deleted"})
}
