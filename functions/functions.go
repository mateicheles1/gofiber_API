package functions

import (
	"golang_api_v2/data"
	"golang_api_v2/models"
	"strconv"

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

	requestBody := new(models.RequestBodyToDo)

	for _, list := range mockData.Data {
		for index, todo := range list.Todos {
			if list.Id == c.Params("listid") && todo.Id == c.Params("todoid") {
				if err := c.BodyParser(requestBody); err != nil {
					return err
				}
				c.JSON(requestBody)
				list.Todos[index].Content = requestBody.Content
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

func CreateToDoByListId(c *fiber.Ctx) error {
	requestBody := new(models.RequestBodyToDo)

	if err := c.BodyParser(requestBody); err != nil {
		return err
	}

	requestBodyCoercion := models.ToDo(*requestBody)

	for index, list := range mockData.Data {
		if list.Id == c.Params("listid") {
			requestBodyCoercion.ListId = c.Params("listid")

			for _, todo := range list.Todos {
				todoid, err := strconv.Atoi(todo.Id)
				if err != nil {
					return err
				}
				requestBodyCoercion.Id = strconv.Itoa(todoid + 1)
			}
			mockData.Data[index].Todos = append(mockData.Data[index].Todos, requestBodyCoercion)
		}
	}

	return nil
}

func GetToDoListById(c *fiber.Ctx) error {
	for _, list := range mockData.Data {
		if list.Id == c.Params("listid") {
			c.JSON(list)
		}
	}

	return nil
}

func UpdateToDoListById(c *fiber.Ctx) error {
	requestBody := new(models.RequestBodyList)

	if err := c.BodyParser(requestBody); err != nil {
		return err
	}

	for index, list := range mockData.Data {
		if list.Id == c.Params("listid") {
			mockData.Data[index].Owner = requestBody.Owner
		}
	}

	return nil
}

func DeleteToDoListById(c *fiber.Ctx) error {

	for index, list := range mockData.Data {
		if list.Id == c.Params("listid") {
			mockData.Data = append(mockData.Data[:index], mockData.Data[index+1:]...)
		}
	}

	return nil
}

func CreateToDoList(c *fiber.Ctx) error {

	requestBodyList := new(models.RequestBodyList)

	if err := c.BodyParser(requestBodyList); err != nil {
		return err
	}

	requestBodyListCoercion := models.ToDoList(*requestBodyList)

	if requestBodyListCoercion.Id < "1" || requestBodyListCoercion.Owner == "" || len(requestBodyListCoercion.Todos) == 0 {
		return fiber.ErrBadRequest
	}

	c.JSON(requestBodyListCoercion)

	mockData.Data = append(mockData.Data, requestBodyListCoercion)

	return nil
}
