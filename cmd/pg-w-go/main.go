package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Country  string
	City     sql.NullString
	Telecode int
}

var schema = `
CREATE TABLE if not exists person (
	first_name text,
	last_name text,
	email text
); 

CREATE TABLE if not exists place (
	country text,
	city text NULL,
	telecode integer
);
`

func main() {
	// connectionString := "user=postgres dbname=gobase password=postgres sslmode=disable"
	connectionString := "postgres://postgres:postgres@localhost:5444/gobase"

	db, err := sqlx.Open("pgx", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	db.MustExec(schema)
	db.MustExec("TRUNCATE person;")
	db.MustExec("TRUNCATE place;")

	tx := db.MustBegin()

	nullStr := sql.NullString{"Test", true}

	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Teest", "Foo", "foo@mail.com")
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Bzz", "Bar", "bar@mail.com")
	tx.NamedExec("INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &Person{"John", "Malkovich", "jhmlkvch@mail.com"})
	tx.MustExec("INSERT INTO place (country, city, telecode) VALUES ($1, $2, $3)", "USA", "New York", "999")
	tx.MustExec("INSERT INTO place (country, telecode) VALUES ($1, $2)", "USA", "429")
	tx.NamedExec("INSERT INTO place (country, city, telecode) VALUES (:country, :city, :telecode)", &Place{"Russia", nullStr, 8800})
	tx.Commit()

	people := []Person{}

	err = db.Select(&people, "SELECT * FROM person ORDER BY last_name DESC")
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range people {
		fmt.Println(p)
	}

	places := []Place{}
	err = db.Select(&places, "SELECT * FROM place ORDER BY country DESC")
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range places {
		fmt.Println(p)
	}

	place := Place{}
	err = db.Get(&place, "SELECT * FROM place WHERE country = $1", "Russia")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Single result:")
	fmt.Println(place)

	person := Person{}
	rows, err := db.Queryx("SELECT * FROM person")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Via scan")
	for rows.Next() {
		rows.StructScan(&person)
		fmt.Println(person)
	}

	_, err = db.NamedQuery("INSERT INTO person (first_name, last_name, email) VALUES (:first, :last, :email)",
		map[string]interface{}{
			"first": "Bin",
			"last":  "Dude",
			"email": "bla@mail.com"})
	if err != nil {
		log.Fatal(err)
	}

	rows, err = db.NamedQuery("SELECT * FROM person WHERE first_name = :first", map[string]interface{}{"first": "Bin"})
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		rows.StructScan(&person)
		fmt.Println("Via map:")
		fmt.Println(person)
	}

	personStructs := []Person{
		{FirstName: "Ardie", LastName: "Savea", Email: "asavea@ab.co.nz"},
		{FirstName: "Sonny Bill", LastName: "Williams", Email: "sbw@ab.co.nz"},
		{FirstName: "Ngani", LastName: "Laumape", Email: "nlaumape@ab.co.nz"},
	}

	_, err = db.NamedExec("INSERT INTO person (first_name, last_name, email) VALUES(:first_name, :last_name, :email)", personStructs)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Get(&person, "SELECT * FROM person WHERE first_name=$1", "Ngani")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Last select:", person)

	schema := []Schema{}

	err = db.Select(&schema, "select column_name, data_type from information_schema.columns where table_name = 'person'")

	if err != nil {
		log.Fatal(err)
	}

	config := getConfig("../../config/db.yaml")
	fmt.Printf("%+v\n", config)

	// var testConfig DatabaseConfig
	// var testTables []Table
	// testTables = []Table{
	// 	Table{
	// 		TableName: "teeest",
	// 		Schemas: []Schema{
	// 			Schema{Name: "id", Type: "text"},
	// 			Schema{Name: "email", Type: "text"},
	// 		},
	// 	},
	// 	Table{
	// 		TableName: "tee_est",
	// 		Schemas: []Schema{
	// 			Schema{Name: "id", Type: "text"},
	// 			Schema{Name: "email", Type: "text"},
	// 		},
	// 	},
	// }

	// testConfig = DatabaseConfig{Tables: testTables}

	// applySchemas(db, testConfig)

}
