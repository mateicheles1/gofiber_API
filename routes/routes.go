package routes

import (
	"golang_api_v2/functions"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleRoutes() {
	app := fiber.New()
	app.Get("api/v2/todoLists", functions.ToDoLists)
	app.Get("api/v2/todos", functions.ToDos)
	app.Get("api/v2/todo/:listid/:todoid", functions.GetToDoById)
	log.Fatal(app.Listen(":8000"))
}
