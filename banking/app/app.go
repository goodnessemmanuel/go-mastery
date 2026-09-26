package app

import (
	"fmt"
	"log"
	"net/http"
)

// function for export must start with capital letter
func Start() {
	//registering the handler or defining the routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getCustomers)

	fmt.Println("starting server at localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
