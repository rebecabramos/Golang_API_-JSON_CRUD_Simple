package main

import (
	"fmt"
	"net/http"
)

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vinda")
}

func configureHandles() {
	http.HandleFunc("/", handleMain) // creates routes in this case the root directory because only the (/) was passed

}

func configureServer() {
	configureHandles()

	fmt.Println("The Server is running on port 1337")
	http.ListenAndServe(":1337", nil) // run the server http (DefaultServerMux)
}

func main() {
	configureServer()
}
