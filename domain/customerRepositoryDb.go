package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hchaudhari73/banking/errs"
	"github.com/hchaudhari73/banking/logger"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	// var rows *sql.Rows
	var err error
	Customers := make([]Customer, 0)

	if status == "" {
		findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		// rows, err = d.client.Query(findAllSQL)
		err = d.client.Select(&Customers, findAllSQL)
	} else {
		findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status=?"
		// rows, err = d.client.Query(findAllSQL, status)
		err = d.client.Select(&Customers, findAllSQL, status)
	}
	if err != nil {
		logger.Error("Error while connectin to database " + err.Error())
		return nil, errs.NewUnexpectedError("Error while connecting to database")
	}

	// err = sqlx.StructScan(rows, &Customers)
	// if err != nil {
	// 	logger.Error("Error while scanning customers " + err.Error())
	// 	return nil, errs.NewUnexpectedError("Error while scanning customers " + err.Error())
	// }

	// for rows.Next() {
	// 	var c Customer
	// 	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.status)
	// 	if err != nil {
	// 		logger.Error("Error while scanning customers " + err.Error())
	// 		return nil, errs.NewUnexpectedError("Error while scanning customers " + err.Error())
	// 	}
	// 	Customers = append(Customers, c)
	// }
	return Customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	// row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := d.client.Get(&c, customerSql, id)
	// err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("User Not Found")
			return nil, errs.NewNotFoundError("User Not Found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
