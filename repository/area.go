package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"shop/contracts"
)

const (
	insertAreaSQL  = `insert into areas (name) values ($1);`
	getAreaSQL     = `select id, name from areas where name=$1;`
	getAllAreasSQL = `select name from areas;`
)

type Areas struct {
	DB *sqlx.DB
}

func (m *Areas) AddArea(name string) error {
	_, err := m.DB.Exec(insertAreaSQL, name)
	if err != nil {
		return errors.Wrap(err, "[AddArea]")
	}

	return nil
}

func (m *Areas) GetArea(name string) (*contracts.Area, error) {
	var area contracts.Area
	err := m.DB.Get(&area, getAreaSQL, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "[GetArea]")
	}

	return &area, nil
}

func (m *Areas) GetAllAreas() ([]string, error) {
	var names []string
	err := m.DB.Select(&names, getAllAreasSQL)
	if err != nil {
		return nil, errors.Wrap(err, "[GetAllAreas]")
	}

	if len(names) == 0 {
		return []string{}, nil
	}

	return names, nil
}
