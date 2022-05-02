package entity

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id       string `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

func (u *User) GetTableName() string {
	return "app_user"
}

func (u *User) Create(db *sqlx.DB) (sql.Result, error) {
	if u.Id != "" {
		u.Id = uuid.New().String()
	}

	res, err :=
		db.NamedExec(
			"INSERT INTO app_user (id, username, password, role) VALUES (:id, :username, :password, :role)",
			u)

	return res, err
}

func (u *User) Read(db *sqlx.DB, id string) error {
	tableName := u.GetTableName()
	err := db.Get(&u, "SELECT * FROM $1 WHERE id=$2", tableName, id)
	return err
}

func (u *User) Update(db *sqlx.DB, id string) error {
	return nil
}

func (u *User) Delete(db *sqlx.DB, id string) (sql.Result, error) {
	return nil, nil
}
