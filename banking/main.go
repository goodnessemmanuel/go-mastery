package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Customer struct represents a customer with their name, city, and zipcode.
// The struct fields are exported (must start with an uppercase letter) to be accessible from other packages.
// but you can specify JSON alias using the `json` tag for the fields
type Customer struct {
	Name    string `json:"firstName"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

func main() {
	//registering the handler or defining the routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getCustomers)

	fmt.Println("starting server at localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, World!")
	if err != nil {
		log.Println(err)
	}
}

// understanding JSON encoding
func getCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"John", "New York", "10001"},
		{"Nelson", "Los Angeles", "90001"},
	}
	//the writer encode header defaults to text/plain text if not specified
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(customers)
	if err != nil {
		log.Println("error encoding customers: ", err)
		return
	}
}
