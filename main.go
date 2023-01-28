package main

import (
	"github.com/1amkaizen/gorestapi_fiber/controllers/bookcontroller"
	"github.com/1amkaizen/gorestapi_fiber/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()

	api := app.Group("/api") // /api
	book := api.Group("/books")

	book.Get("/", bookcontroller.Index)
	book.Get("/:id", bookcontroller.Show)
	book.Post("/", bookcontroller.Create)
	book.Put("/:id", bookcontroller.Update)
	book.Delete("/:id", bookcontroller.Delete)

	app.Listen(":3000")

}
