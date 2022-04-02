SHELL = bash

PGPORT     ?= 5444
PGHOST     ?= localhost
PGUSER     ?= postgres
PGDATABASE ?= gobase
PGPASSWORD ?= postgres
PGIMAGE    ?= postgres:latest

.EXPORT_ALL_VARIABLES:

run:
	cd cmd/pg-w-go && go run *.go
up:
	docker-compose up -d
down:
	docker-compose down