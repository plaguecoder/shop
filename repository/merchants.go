package repository

import (
	"database/sql"
	"github.com/pkg/errors"
)

const (
	insertMerchantSQL = `insert into merchants (area, name, phone) values ($1, $2, $3);`
)

type Merchants struct {
	DB *sql.DB
}

func (m *Merchants) AddMerchant(area, name, phone string) error {
	_, err := m.DB.Exec(insertMerchantSQL, area, name, phone)
	if err != nil {
		return errors.Wrap(err, "fata")
	}

	return nil
}
