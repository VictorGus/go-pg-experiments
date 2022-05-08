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
	if u.Id == "" {
		u.Id = uuid.New().String()
	}

	res, err :=
		db.NamedExec(
			"INSERT INTO app_user (id, username, password, role) VALUES (:id, :username, :password, :role)",
			u)

	return res, err
}

func (u *User) Read(db *sqlx.DB, id string) error {
	// tableName := u.GetTableName()
	err := db.Get(u, "SELECT * FROM app_user WHERE id = $1", id)
	return err
}

func (u *User) Update(db *sqlx.DB) (sql.Result, error) {
	res, err :=
		db.NamedExec(
			"UPDATE app_user SET username = :username, password = :password, role = :role WHERE id = :id",
			u)

	return res, err
}

func (u *User) Delete(db *sqlx.DB, id string) (sql.Result, error) {
	res, err :=
		db.Exec("DELETE FROM app_user WHERE id = $1", id)
	return res, err
}
