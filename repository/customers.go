package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"shop/contracts"
)

const (
	insertCustomerSQL   = `insert into customers (name, area_id, phone, description) values ($1, $2, $3, $4);`
	getAllCustomersSQL  = `select c.id, c.name, a.name as area, c.phone, c.description from customers c inner join areas a on c.area_id = a.id`
	getCustomersByIDSQL = `select c.id, c.name, a.name as area, c.phone, c.description from customers c inner join areas a on c.area_id = a.id where c.id = $1;`
)

type Customers struct {
	DB *sqlx.DB
}

func (m *Customers) AddCustomer(customer *contracts.Customer) error {
	_, err := m.DB.Exec(insertCustomerSQL, customer.Name, customer.AreaID, customer.Phone, customer.Description)
	if err != nil {
		return errors.Wrap(err, "[AddCustomer]")
	}

	return nil
}

func (m *Customers) GetCustomer(id int64) (*contracts.Customer, error) {
	var customer contracts.Customer
	err := m.DB.Get(&customer, getCustomersByIDSQL, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "[GetCustomer]")
	}

	return &customer, nil
}

func (m *Customers) GetAllCustomers() ([]contracts.Customer, error) {
	var customers []contracts.Customer
	err := m.DB.Select(&customers, getAllCustomersSQL)
	if err != nil {
		return nil, errors.Wrap(err, "[GetAllCustomers]")
	}

	if len(customers) == 0 {
		return []contracts.Customer{}, nil
	}

	return customers, nil
}
