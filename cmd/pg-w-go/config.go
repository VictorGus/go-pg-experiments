package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
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

func getConfig(configPath string) DatabaseConfig {
	dataBaseConfig := DatabaseConfig{}

	err := readFile(&dataBaseConfig, configPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return dataBaseConfig

}

func readFile(cfg *DatabaseConfig, pth string) error {
	f, err := os.Open(pth)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	} else {
		return nil
	}
}
