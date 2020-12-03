package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/", homepage).Methods("GET")
	myRouter.HandleFunc("/customers", getAllCustomers).Methods("GET")
	myRouter.HandleFunc("/customer/{customer_id:[0-9]+}", getCustomer).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
