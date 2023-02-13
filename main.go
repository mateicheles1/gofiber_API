package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Item struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

var items []Item

func sayHello(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	io.WriteString(w, "Welcome to the homepage :)")
}

func getItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

func getOneItem(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	for _, todo := range items {
		if todo.Id == id {
			json.NewEncoder(w).Encode(todo)
		}
	}

}

func createItem(w http.ResponseWriter, r *http.Request) {
	var toDo Item

	if r.Method != "POST" {
		fmt.Fprintln(w, "wrong method, needs POST")
		return
	}

	err := json.NewDecoder(r.Body).Decode(&toDo)
	if err != nil {
		log.Println("error reading request body", err)
		return
	}

	if toDo.Id == 0 || toDo.Content == "" {
		fmt.Fprintln(w, "error writing todo; content or id invalid")
		return
	}

	items = append(items, toDo)
	json.NewEncoder(w).Encode(toDo)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {

	if r.Method != "DELETE" {
		fmt.Fprintln(w, "wrong method, needs DELETE")
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		return
	}

	for index, todo := range items {
		if todo.Id == id {
			items = append(items[:index], items[index+1:]...)
		}
	}

	json.NewEncoder(w).Encode(items)
}

func patchItem(w http.ResponseWriter, r *http.Request) {
	var toDo Item

	if r.Method != "PATCH" {
		fmt.Fprintln(w, "wrong method, needs PATCH")
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	json.NewDecoder(r.Body).Decode(&toDo)

	if toDo.Content == "" {
		fmt.Fprintln(w, "error patching todo; invalid content")
	}

	for index, todoItem := range items {
		if todoItem.Id == id {
			todoItem.Content = toDo.Content
			items[index] = todoItem
		}
	}
	json.NewEncoder(w).Encode(toDo)
}
func handleRoutes() {
	http.HandleFunc("/item/create", createItem)
	http.HandleFunc("/item/delete", deleteItem)
	http.HandleFunc("/item/patch", patchItem)
	http.HandleFunc("/item/read", getOneItem)
	http.HandleFunc("/items", getItems)
	http.HandleFunc("/", sayHello)
	log.Fatal(http.ListenAndServe(":1000", nil))
}

func main() {
	items = []Item{
		{1, "Washing the dishes."},
		{2, "Learning RESTful API implementation."},
		{3, "Watch GO Documentation."},
	}

	handleRoutes()
}
