package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

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

func EnrichWithCreds(config *DatabaseConfig) {
	config.DBUser = os.Getenv("PGUSER")
	config.DBPassword = os.Getenv("PGPASSWORD")
	config.DBHost = os.Getenv("PGHOST")
	config.DBPort = os.Getenv("PGPORT")
	config.DBName = os.Getenv("PGDATABASE")
}

func InitDataBase(config DatabaseConfig) {
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
		db.MustExec("CREATE DATABASE " + config.DBName)
		log.Printf("%s db has been initiated\n", config.DBName)
	} else {
		log.Printf("%s db already exists", config.DBName)
	}

	db.Close()
	appDB, err := sqlx.Open("pgx", connectionString+"/"+config.DBName)
	if err != nil {
		log.Fatal(err)
	}

	applySchemas(appDB, config)
	appDB.Close()
	db, err = sqlx.Open("pgx", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	queryString = fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s_test'", config.DBName)
	rows, err = db.Queryx(queryString)
	if err != nil {
		log.Fatal(err)
	}

	if !rows.Next() {
		db.MustExec("CREATE DATABASE " + config.DBName + "_test TEMPLATE " + config.DBName)
		log.Printf("%s_test db has been initiated\n", config.DBName)
	} else {
		log.Printf("%s_test db already exists", config.DBName)
	}

	defer db.Close()
}
