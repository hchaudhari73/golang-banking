package domain

import (
	"log"

	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hchaudhari73/banking/errs"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {

	findAllSql := "select * from customers"
	rows, err := d.client.Query(findAllSql)
	if err != nil {
		return nil, errs.NewUnexpectedError("Error while connecting to database")
	}

	Customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			return nil, errs.NewUnexpectedError("Error while scanning customer")
		}
		Customers = append(Customers, c)
	}
	return Customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("User Not Found")
		} else {
			log.Printf("Error while scannin customer %b\n", err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:mbatman@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
