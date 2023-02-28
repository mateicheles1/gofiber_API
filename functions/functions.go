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
			}
		}
	}

	return nil
}

func UpdateToDoById(c *fiber.Ctx) error {

	updatedContent := new(models.UpdatedContent)

	for _, list := range mockData.Data {
		for index, todo := range list.Todos {
			if list.Id == c.Params("listid") && todo.Id == c.Params("todoid") {
				if err := c.BodyParser(updatedContent); err != nil {
					return err
				}
				c.JSON(updatedContent)
				list.Todos[index].Content = updatedContent.Content
			}
		}
	}

	return nil

}

func DeleteToDoById(c *fiber.Ctx) error {

	for i := 0; i < len(mockData.Data); i++ {
		for j := 0; j < len(mockData.Data[i].Todos); j++ {
			if mockData.Data[i].Id == c.Params("listid") && mockData.Data[i].Todos[j].Id == c.Params("todoid") {
				mockData.Data[i].Todos = append(mockData.Data[i].Todos[:j], mockData.Data[i].Todos[j+1:]...)
			}
		}
	}

	return nil
}
