package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bem vinda")
	}) // creates routes in this case the root directory because only the (/) was passed

	fmt.Println("The Server is running on port 1337")
	http.ListenAndServe(":1337", nil) // run the server http (DefaultServerMux)
}
