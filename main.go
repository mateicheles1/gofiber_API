package main

import (
	"fmt"
	"golang_api_v2/models"
)

var data models.AppData

func main() {
	data = models.AppData{
		Data: []models.ToDoList{
			{Id: "1", Owner: "Matei Cheles", Todos: []models.ToDo{
				{Id: "1", ListId: "1", Content: "Do the dishes"},
				{Id: "2", ListId: "1", Content: "Go for a walk"},
				{Id: "3", ListId: "1", Content: "Grab a coffee"},
			},
			},
			{Id: "2", Owner: "Tudor Datcu", Todos: []models.ToDo{
				{Id: "1", ListId: "2", Content: "Watch a docuseries"},
				{Id: "2", ListId: "2", Content: "Have daily"},
				{Id: "3", ListId: "2", Content: "Check api v2"},
			},
			},
		},
	}

	fmt.Println(data)
}
