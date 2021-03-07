package domain

import (
	"strconv"

	"github.com/hchaudhari73/banking/errs"
	"github.com/hchaudhari73/banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {

	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (? ? ? ? ?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

/*
	For creating a user
*/

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
