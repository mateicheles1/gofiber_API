package functions

import (
	"golang_api_v2/data"
	"golang_api_v2/models"

	"github.com/gofiber/fiber/v2"
)

var mockData = data.MockData

func ToDoLists(c *fiber.Ctx) error {

	c.JSON(mockData.Data)

	return nil
}

func ToDos(c *fiber.Ctx) error {

	var todos []models.ToDo

	for _, list := range mockData.Data {
		todos = append(todos, list.Todos...)
	}

	c.JSON(todos)

	return nil
}

func GetToDoById(c *fiber.Ctx) error {

	for _, list := range mockData.Data {
		for _, todo := range list.Todos {
			if todo.ListId == c.Params("listid") && todo.Id == c.Params("todoid") {
				c.JSON(todo)
			} else {
				return fiber.ErrBadRequest
			}
		}
	}

	return nil
}

// func UpdateToDoById(c *fiber.Ctx) error {

// }
