package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hchaudhari73/banking/domain"
	"github.com/hchaudhari73/banking/service"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variables not defined...")
	}
}

func Start() {

	sanityCheck()

	address := os.Getenv("SERVER_PORT")
	fmt.Printf("Starting server at port %s...", address)
	myRouter := mux.NewRouter()

	//wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	myRouter.HandleFunc("/customers", ch.getAllCustomers).Methods("GET")
	myRouter.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods("GET")

	address = os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), myRouter))
}
