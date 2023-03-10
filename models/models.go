package models

type ToDo struct {
	Id      string `json:"id"`
	ListId  string `json:"listid"`
	Content string `json:"content"`
}

type ToDoList struct {
	Id    string `json:"listId"`
	Owner string `json:"owner"`
	Todos []ToDo `json:"todos"`
}

type AppData struct {
	Data []ToDoList `json:"data"`
}

type RequestBodyToDo struct {
	Id      string `json:"id"`
	ListId  string `json:"listidrequest"`
	Content string `json:"content"`
}

type RequestBodyList struct {
	Id    string `json:"listId"`
	Owner string `json:"owner"`
	Todos []ToDo `json:"todos"`
}
