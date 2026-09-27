package app

import (
	"fmt"
	"log"
	"net/http"
)

// Start function for export must start with a capital letter
func Start() {
	startWithCustomMultiplexer()
}

func startWithCustomMultiplexer() {
	//defining own mux handler to register the routes
	mux := http.NewServeMux()

	//registering routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getCustomers)

	fmt.Println("starting server at localhost:8000 with custom mux")
	log.Fatal(http.ListenAndServe(":8000", mux))
}

func startWithDefaultHttpMultiplexer() {
	//registering the routes to be handled by the default http.Server Mux
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getCustomers)

	fmt.Println("starting server at localhost:8000 with default http.Server Mux")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
