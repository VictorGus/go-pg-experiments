package entity

import (
	"reflect"
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

	if intRes != 1 {
		t.Logf("Expected:\n%v\nGot:\n%v", 1, intRes)
		t.Fail()
	}

	db.MustExec("TRUNCATE app_user;")

	userWithoutID := &User{Username: "Test", Password: "foobar", Role: "admin"}
	res, err = userWithoutID.Create(db)
	if err != nil {
		t.Error(err)
	}

	intRes, err = res.RowsAffected()
	if err != nil {
		t.Error(err)
	}

	if intRes != 1 {
		t.Logf("Expected:\n%v\nGot:\n%v", 1, intRes)
		t.Fail()
	}

	db.MustExec("TRUNCATE app_user;")

	defer db.Close()
}

func TestRead(t *testing.T) {
	user := &User{"123", "Test", "foobar", "admin"}

	db, err := sqlx.Open("pgx", ConnectionString)
	if err != nil {
		t.Error(err)
	}

	_, err = user.Create(db)
	if err != nil {
		t.Error(err)
	}

	userFromDB := &User{}
	err = userFromDB.Read(db, "123")
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(user, userFromDB) {
		t.Logf("Expected:\n%v\nGot:\n%v", user, userFromDB)
		t.Fail()
	}

	db.MustExec("TRUNCATE app_user;")
	defer db.Close()
}

func TestUpdate(t *testing.T) {
	user := &User{"123", "Test", "foobar", "admin"}

	db, err := sqlx.Open("pgx", ConnectionString)
	if err != nil {
		t.Error(err)
	}

	_, err = user.Create(db)
	if err != nil {
		t.Error(err)
	}

	updatedUser := &User{"123", "Foo", "Bar", "manager"}
	res, err := updatedUser.Update(db)
	if err != nil {
		t.Error(err)
	}

	intRes, err := res.RowsAffected()
	if err != nil {
		t.Error(err)
	}

	if intRes != 1 {
		t.Logf("Expected:\n%v\nGot:\n%v", 1, intRes)
		t.Fail()
	}

	err = user.Read(db, "123")
	if err != nil {
		t.Error(err)
	}

	db.MustExec("TRUNCATE app_user;")
	defer db.Close()
}

func TestDelete(t *testing.T) {
	user := &User{"123", "Test", "foobar", "admin"}

	db, err := sqlx.Open("pgx", ConnectionString)
	if err != nil {
		t.Error(err)
	}

	_, err = user.Create(db)
	if err != nil {
		t.Error(err)
	}

	res, err := user.Delete(db, "123")
	if err != nil {
		t.Error(err)
	}

	intRes, err := res.RowsAffected()

	if err != nil {
		t.Error(err)
	}

	if intRes != 1 {
		t.Logf("Expected:\n%v\nGot:\n%v", 1, intRes)
		t.Fail()
	}

	db.MustExec("TRUNCATE app_user;")
	defer db.Close()
}
