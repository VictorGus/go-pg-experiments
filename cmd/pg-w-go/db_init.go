package main

import (
	"fmt"
	"log"
	"strings"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Schema struct {
	Name string `db:"column_name"`
	Type string `db:"data_type"`
}

type Table struct {
	TableName string `yaml:"tableName"`
	Schemas   []Schema
}

type DatabaseConfig struct {
	Tables     []Table
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
}

func applySchemas(db *sqlx.DB, config DatabaseConfig) {
	tables := config.Tables
	for _, table := range tables {
		var gluedParts []string
		for _, schema := range table.Schemas {
			gluedParts = append(gluedParts, schema.Name+" "+schema.Type)
		}
		gluedPartsString := strings.Join(gluedParts, ",")
		db.MustExec("CREATE TABLE if not exists " + table.TableName + "(" + gluedPartsString + ")")
	}
}

func InitDataBase(config DatabaseConfig) {
	// connectionString := "postgres://postgres:postgres@localhost:5444"
	connectionString := fmt.Sprintf(
		"%s://%s:%s@%s:%s", config.DBUser, config.DBUser, config.DBPassword, config.DBHost, config.DBPort)
	db, err := sqlx.Open("pgx", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	queryString := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s'", config.DBName)

	rows, err := db.Queryx(queryString)
	if err != nil {
		log.Fatal(err)
	}

	if !rows.Next() {
		db.MustExec("CREATE DATABASE gobase_test TEMPLATE gobase")
		log.Println("Test db has been initiated")
	} else {
		log.Println("Test db already exists")
	}

	defer db.Close()
}
