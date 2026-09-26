package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//registering the handler. defining the routes
	http.HandleFunc("/greet", greet)
	fmt.Println("starting server at localhost:8000")

	//starting the server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, World!")
	if err != nil {
		log.Println(err)
	}
}
