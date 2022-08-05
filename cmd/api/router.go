package main

import "github.com/gofiber/fiber/v2"

func (app *application) Router() *fiber.App {
	r := fiber.New()
	r.Post("/api/v1/things", app.InsertThing)
	r.Get("/api/v1/things/:id", app.ReadThing)
	r.Get("/api/v1/things", app.FetchThing)
	r.Delete("/api/v1/things/:id", app.RemoveThing)
	return r
}
