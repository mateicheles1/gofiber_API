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

	// Am luat id-ul din URL

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// Daca id-ul din URL este acelasi cu todo id-ul serv raspunde cu todo-ul respectiv

	for _, todo := range items {
		if todo.Id == id {
			json.NewEncoder(w).Encode(todo)
		}
	}

}

func createItem(w http.ResponseWriter, r *http.Request) {

	// am creat o variabila toDo de tip Item care va fi populata cu request body
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

	// adaug noul todo la existentele todo-uri si dau ca si raspuns todo-ul adaugat

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
	// daca conditia e true, tai din array-ul respectiv todo-ul in cauza
	for index, todo := range items {
		if todo.Id == id {
			items = append(items[:index], items[index+1:]...)
		}
	}

	json.NewEncoder(w).Encode(items)
}

func patchItem(w http.ResponseWriter, r *http.Request) {

	// la fel, creez o variabila care va fi populata cu request body
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

	// probabil nu e cea mai ok metoda de a o face, dar am zis ca daca id-ul coincide cu id-ul copiei todo-ului, todo-ului i se va schimba contentul cu contentul variabilei care a fost populata cu corpul requestului, iar todo-ul care se afla la index-ul in cauza va prelua datele copiei.

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
