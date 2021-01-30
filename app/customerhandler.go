package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hchaudhari73/banking/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: homepage")
	fmt.Fprintf(w, "Welcome to homepage")
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	// customers := []Customer{
	// 	{"Harshal", "Dombivli", "421202"},
	// 	{"Bruce", "Gautom", "67709"},
	// }

	fmt.Println("Endpoint hit: getAllCustomers")

	customers, _ := ch.service.GetAllCustomer(status)
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getCustomer")

	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
