package main

import (
	"fmt"
	"strings"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Schema struct {
	Name string `db:"column_name"`
	Type string `db:"data_type"`
}

type Table struct {
	TableName string
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
		res := db.MustExec("CREATE TABLE if not exists " + table.TableName + "(" + gluedPartsString + ")")
		fmt.Println(res)
	}
}
