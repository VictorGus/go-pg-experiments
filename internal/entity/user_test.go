package entity

import (
	"testing"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

var ConnectionString string = "postgres://postgres:postgres@localhost:5444/gobase_test"

func TestCreate(t *testing.T) {
	user := &User{"123", "Test", "foobar", "admin"}

	db, err := sqlx.Open("pgx", ConnectionString)
	if err != nil {
		t.Error(err)
	}

	res, err := user.Create(db)
	if err != nil {
		t.Error(err)
	}

	intRes, err := res.RowsAffected()
	if err != nil {
		t.Error(err)
	}

	if 1 != intRes {
		t.Logf("Expected:\n%v\nGot:\n%v", 1, intRes)
		t.Fail()
	}
}
