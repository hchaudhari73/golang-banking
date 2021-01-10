package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hchaudhari73/banking/domain"
	"github.com/hchaudhari73/banking/service"
)

func Start() {
	fmt.Println("Starting server at port 8080...")
	myRouter := mux.NewRouter()

	//wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	myRouter.HandleFunc("/customers", ch.getAllCustomers).Methods("GET")
	myRouter.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
