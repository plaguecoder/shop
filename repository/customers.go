package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"shop/contracts"
)

const (
	insertCustomerSQL  = `insert into customers (name, area, phone, description) values ($1, $2, $3, $4);`
	getAllCustomersSQL = `select id, name, area, phone, description from customers;`
)

type Customers struct {
	DB *sqlx.DB
}

func (m *Customers) AddCustomer(customer *contracts.Customer) error {
	_, err := m.DB.Exec(insertCustomerSQL, customer.Name, customer.Area, customer.Phone, customer.Description)
	if err != nil {
		return errors.Wrap(err, "[AddCustomer]")
	}

	return nil
}

func (m *Customers) GetAllCustomers() ([]contracts.Customer, error) {
	var customers []contracts.Customer
	err := m.DB.Select(&customers, getAllCustomersSQL)
	if err != nil {
		return nil, errors.Wrap(err, "[GetAllCustomers]")
	}

	return customers, nil
}
