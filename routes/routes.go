package routes

import (
	"golang_api_v2/data"
	"golang_api_v2/functions"
	"log"

	"github.com/gofiber/fiber/v2"
)

var mockData = data.MockData

func HandleRoutes() {
	app := fiber.New()
	app.Get("/", functions.HomePage)
	log.Fatal(app.Listen(":8000"))
}
