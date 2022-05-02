package operations

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func Create(db *sqlx.DB, entity string, data interface{}) (sql.Result, error) {
	res, err := db.NamedExec(
		"INSERT INTO person (id, name, email, address) VALUES (:id, :name, :email, address)",
		data) // shaper of params depending on resourceType

	return res, err
}

func Read(db *sqlx.DB, entity string, id string) (sql.Result, error) {
	return nil, nil
}

func Update(db *sqlx.DB, entity string, id string, data interface{}) (sql.Result, error) {
	return nil, nil
}

func Delete(db *sqlx.DB, entity string, id string) (sql.Result, error) {
	return nil, nil
}
