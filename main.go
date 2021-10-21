package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct {
	Id     int
	Title  string
	Author string
}

var Books []Book = []Book{
	Book{
		Id:     1,
		Title:  "O Guarani",
		Author: "José de Alencar",
	},
	Book{
		Id:     2,
		Title:  "Cazuza",
		Author: "Viriato Correia",
	},
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vinda")
}

func listBooks(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w) // returns a new encoder that writes to w
	encoder.Encode(Books)
}

func configureHandles() {
	http.HandleFunc("/", handleMain) // creates routes in this case the root directory because only the (/) was passed
	http.HandleFunc("/Books", listBooks)
}

func configureServer() {
	configureHandles()

	fmt.Println("The Server is running on port 1337")
	http.ListenAndServe(":1337", nil) // run the server http (DefaultServerMux)
}

func main() {
	configureServer()
}
