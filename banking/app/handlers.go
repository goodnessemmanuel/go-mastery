package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

// Customer struct represents a customer with their name, city, and zipcode.
// The struct fields are exported (must start with an uppercase letter) to be accessible from other packages.
// but you can specify JSON alias using the `json` tag for the fields
type Customer struct {
	Name    string `json:"firstName" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, World!")
	if err != nil {
		log.Println(err)
	}
}

// understanding JSON/XML encoding
func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"John", "New York", "10001"},
		{"Nelson", "Los Angeles", "90001"},
	}

	reqHeaderContentType := r.Header.Get("Content-Type")

	//the writer encode header defaults to text/plain text if not specified
	if reqHeaderContentType == "application/xml" {
		// xml encoding
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		//json encoding
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
