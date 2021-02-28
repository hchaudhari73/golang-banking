package domain

import (
	"github.com/hchaudhari73/banking/dto"
	"github.com/hchaudhari73/banking/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string `db:"name"`
	City        string `db:"city"`
	Zipcode     string `db:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `db:"status"`
}

/*
	To return the data in the form of dto,
	we can write a method to get the response on dto
	Instead of doing it in the service method.
*/

func (c Customer) statusAsText() string {
	if c.Status == "0" {
		return "inactive"
	}
	return "active"
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
