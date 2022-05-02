package entity

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Entitier interface {
	Create(*sqlx.DB) (sql.Result, error)
	Read(*sqlx.DB, string) (Entitier, error)
	Update(*sqlx.DB, string) (Entitier, error)
	Delete(*sqlx.DB, string) (sql.Result, error)
	GetTableName() string
}
