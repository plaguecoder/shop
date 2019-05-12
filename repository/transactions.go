package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"shop/contracts"
)

const (
	insertTransactionSQL           = `insert into transactions (customer_id, date, amount, type, description) values ($1, $2, $3, $4, $5);`
	getTransactionsByCustomerIDSQL = `select id, customer_id, date, amount, type, description from transactions where customer_id = $1 order by date desc;`
)

type Transactions struct {
	DB *sqlx.DB
}

func (m *Transactions) AddTransaction(t *contracts.Transaction) error {
	_, err := m.DB.Exec(insertTransactionSQL, t.CustomerID, t.Date, t.Amount, t.Type, t.Description)
	if err != nil {
		return errors.Wrap(err, "[AddTransaction]")
	}

	return nil
}

func (m *Transactions) GetAllTransactions(customerID int64) ([]contracts.Transaction, error) {
	var transactions []contracts.Transaction
	err := m.DB.Select(&transactions, getTransactionsByCustomerIDSQL, customerID)
	if err != nil {
		return nil, errors.Wrap(err, "[GetAllTransactions]")
	}

	if len(transactions) == 0 {
		return []contracts.Transaction{}, nil
	}

	return transactions, nil
}
