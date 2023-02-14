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

func returnQueryId(w http.ResponseWriter, r *http.Request) int {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		io.WriteString(w, "Invalid ID\n")
		http.NotFound(w, r)
		return 0
	}
	return id
}

func showLists(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		io.WriteString(w, "Wrong methods; needs GET")
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(lists)
}

func showOneList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		io.WriteString(w, "Wrong method; needs GET")
		http.NotFound(w, r)
		return
	}
	id := returnQueryId(w, r)
	for _, list := range lists {
		if list.Id == id {
			json.NewEncoder(w).Encode(list)
		}
	}
}

func showOneTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		io.WriteString(w, "Wrong method; needs GET")
		http.NotFound(w, r)
		return
	}

	id := returnQueryId(w, r)
	secondId, err := strconv.Atoi(r.URL.Query().Get("secId"))
	if err != nil || secondId < 1 {
		io.WriteString(w, "invalid todo id\n")
		http.NotFound(w, r)
		return

	}
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

func handleRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		io.WriteString(w, "Welcome tot the homepage")
	})
	http.HandleFunc("/api/v1/lists", showLists)
	http.HandleFunc("/api/v1/list", showOneList)
	http.HandleFunc("/api/v1/lists/todos", showOneTodo)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	lists = []todoList{
		{
			Id:    1,
			Owner: "Matei Cheles",
			Todos: []todoItem{
				{1, "washing the dishes"},
				{2, "watching netflix"},
				{3, "going out"},
			},
		},
		{
			Id:    2,
			Owner: "Tudor Datcu",
			Todos: []todoItem{
				{1, "water the plants"},
				{2, "check matei's work"},
				{3, "watch a docuseries"},
			},
		},
	}

	handleRoutes()
}
