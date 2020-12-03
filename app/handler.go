package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: homepage")
	fmt.Fprintf(w, "Welcome to homepage")
}

var customers = []Customer{
	{
		Name:    "Harshal",
		City:    "Dombivli",
		Zipcode: "421202",
	},
	{
		Name:    "Bruce",
		City:    "Goutam",
		Zipcode: "45202",
	},
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getCustomers")
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getCustomer")
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}
