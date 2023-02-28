package functions

import (
	"golang_api_v2/data"
	"golang_api_v2/models"

	"github.com/gofiber/fiber/v2"
)

var mockData = data.MockData

func ToDoLists(c *fiber.Ctx) error {
	return c.JSON(mockData)
}

func ToDos(c *fiber.Ctx) error {
	var todos []models.ToDo

	for _, list := range mockData.Data {
		todos = append(todos, list.Todos...)
	}

	return c.JSON(todos)
}
