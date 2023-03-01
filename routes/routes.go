package routes

import (
	"golang_api_v2/functions"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleRoutes() {
	app := fiber.New()

	app.Get("api/v2/todolists", functions.ToDoLists)
	app.Get("api/v2/todos", functions.ToDos)

	app.Get("api/v2/todo/:listid/:todoid", functions.GetToDoById)
	app.Patch("api/v2/todo/:listid/:todoid", functions.UpdateToDoById)
	app.Delete("api/v2/todo/:listid/:todoid", functions.DeleteToDoById)
	app.Post("api/v2/todo/:listid", functions.CreateToDoByListId)

	app.Get("api/v2/todolist/:listid", functions.GetToDoListById)
	app.Patch("api/v2/todolist/:listid", functions.UpdateToDoListById)
	app.Delete("api/v2/todolist/:listid", functions.DeleteToDoListById)
	app.Post("api/v2/todolist/", functions.CreateToDoList)

	log.Fatal(app.Listen(":8000"))
}
