package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Start function for export must start with a capital letter
func Start() {
	startWithGorillaMux()
}

func startWithCustomMultiplexer() {
	//defining own mux handler to register the routes
	mux := http.NewServeMux()

	//registering routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	fmt.Println("starting server at localhost:8000 with custom mux")
	log.Fatal(http.ListenAndServe(":8000", mux))
}

func startWithDefaultHttpMultiplexer() {
	//registering the routes to be handled by the default http.Server Mux
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	fmt.Println("starting server at localhost:8000 with default http.Server Mux")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

/**
 * using gorilla mux to register the routes
 * gorilla mux is a library that provides a high-performance, idiomatic and extensible HTTP router for Go
 */
func startWithGorillaMux() {
	//using gorilla mux to register the routes
	router := mux.NewRouter()
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers)
	//only numbers are allowed. using gorilla mux regex to validate the id in the route path
	// entering a non-numeric value will auto result in a 404 page not found
	router.HandleFunc("/customers/{id:[0-9]+}", getCustomer)

	fmt.Println("starting server at localhost:8000 with gorilla mux")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getCustomer(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	customerId := vars["id"]
	fmt.Fprintf(writer, "Customer ID: %s", customerId)
}
