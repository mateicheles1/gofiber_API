package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type todoItem struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

type todoList struct {
	Id    int        `json:"id"`
	Owner string     `json:"owner"`
	Todos []todoItem `json:"todos"`
}

var lists []todoList

/*

########## RETURNING THE URL IDS FOR THE LISTS/TODOS ##########

*/

func returnQueryId(w http.ResponseWriter, r *http.Request) int {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		io.WriteString(w, "invalid list id\n")
		http.NotFound(w, r)
		return 0
	}
	return id
}

func returnSecondQueryId(w http.ResponseWriter, r *http.Request) int {
	id, err := strconv.Atoi(r.URL.Query().Get("secId"))
	if err != nil || id < 1 {
		io.WriteString(w, "invalid todo id\n")
		http.NotFound(w, r)
		return 0
	}
	return id
}

/*
########## CRUD for the LISTS ##########
*/

func showLists(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		io.WriteString(w, "Wrong methods; needs GET")
		return
	}
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	io.WriteString(w, "These are all the lists\n")
	json.NewEncoder(w).Encode(lists)
}

func showOneList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		io.WriteString(w, "Wrong method; needs GET")
		return
	}
	id := returnQueryId(w, r)
	for _, list := range lists {
		if list.Id == id {
			json.NewEncoder(w).Encode(list)
		}
	}
}

func createList(w http.ResponseWriter, r *http.Request) {

	var list todoList

	if r.Method != "POST" {
		io.WriteString(w, "wrong method, needs POST")
		return
	}

	json.NewDecoder(r.Body).Decode(&list)
	if list.Id < 1 || list.Owner == "" || len(list.Todos) == 0 {
		io.WriteString(w, "error writing list")
		return
	}
	lists = append(lists, list)
	json.NewEncoder(w).Encode(list)

}

func updateList(w http.ResponseWriter, r *http.Request) {

	var placeHolderList todoList

	if r.Method != "PATCH" {
		io.WriteString(w, "wrong method; needs PATCH")
		return
	}
	id := returnQueryId(w, r)
	json.NewDecoder(r.Body).Decode(&placeHolderList)

	for index, list := range lists {
		if list.Id == id {
			list.Owner = placeHolderList.Owner
			lists[index] = list
		}
	}
	json.NewEncoder(w).Encode(lists)
}

func deleteList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		io.WriteString(w, "wrong method; needs DELETE")
	}

	id := returnQueryId(w, r)

	for i, list := range lists {
		if list.Id == id {
			json.NewEncoder(w).Encode(list)
			lists = append(lists[:i], lists[i+1:]...)
		}
	}
}

/*
########## CRUD for the TODOS ##########
*/

func showOneTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		io.WriteString(w, "Wrong method; needs GET")
		http.NotFound(w, r)
		return
	}

	id := returnQueryId(w, r)
	secondId := returnSecondQueryId(w, r)
	for _, list := range lists {
		if list.Id == id {
			for _, todo := range list.Todos {
				if todo.Id == secondId {
					json.NewEncoder(w).Encode(todo)
				}
			}
		}
	}

}

func createOneTodo(w http.ResponseWriter, r *http.Request) {

	var todo todoItem
	if r.Method != "POST" {
		io.WriteString(w, "wrong method; needs POST")
		return
	}
	id := returnQueryId(w, r)

	json.NewDecoder(r.Body).Decode(&todo)
	if todo.Content == "" || todo.Id == 0 {
		io.WriteString(w, "invalid writing todo")
		return
	}

	for index, list := range lists {
		if list.Id == id {
			lists[index].Todos = append(lists[index].Todos, todo)
			json.NewEncoder(w).Encode(lists[index])
		}
	}
}

/*
########## ROUTES ##########
*/

func handleRoutes() {
	http.HandleFunc("/api/v1/lists/todo/create", createOneTodo)
	http.HandleFunc("/api/v1/lists/todos", showOneTodo)
	http.HandleFunc("/", showLists)
	http.HandleFunc("/api/v1/list", showOneList)
	http.HandleFunc("/api/v1/list/create", createList)
	http.HandleFunc("/api/v1/list/update", updateList)
	http.HandleFunc("/api/v1/list/delete", deleteList)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

/*
########## MAIN ##########
*/

func main() {
	lists = []todoList{
		{
			Id:    1,
			Owner: "Matei Cheles",
			Todos: []todoItem{
				{1, "learn Go best practices"},
				{2, "watch netflix"},
				{3, "learn RESTful"},
			},
		},
		{
			Id:    2,
			Owner: "Tudor Datcu",
			Todos: []todoItem{
				{1, "grab a coffee"},
				{2, "check matei's work"},
				{3, "watch a docuseries"},
			},
		},
	}

	handleRoutes()
}
