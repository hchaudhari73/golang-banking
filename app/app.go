package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/hchaudhari73/banking/domain"
	"github.com/hchaudhari73/banking/service"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variables not defined...")
	}
}

/*
	Instead of creating connection pool in repositories
	we are creating connection in wiring section.
*/
func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddr, dbPort, dbName)
	// dataSource := fmt.Sprintf("%s@tcp(%s:%s)/%s", dbUser, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

func Start() {

	sanityCheck()

	address := os.Getenv("SERVER_PORT")
	fmt.Printf("Starting server at port %s...", address)
	myRouter := mux.NewRouter()

	/*
		wiring function for creating new customer
	*/
	dbClient := getDbClient()

	//wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	// accountRepositryDb := domain.NewAccountRepositoryDb(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}

	myRouter.HandleFunc("/customers", ch.getAllCustomers).Methods("GET")
	myRouter.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods("GET")

	address = os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), myRouter))
}
