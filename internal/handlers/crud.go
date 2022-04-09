package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func cutOutEntity(url string) string {
	return strings.Split(url, "/")[0]
}

func formValuesSeq(url string) string {
	valueSequences := map[string][]string{
		"person":       {"id", "name", "email", "address"},
		"organization": {"id", "name", "email", "password"},
	}

	entityName := cutOutEntity(url)
	return strings.Join(valueSequences[entityName], ", ")
}

func HandleCreate(ctx echo.Context) error {
	connection := ctx.Get("db-connection").(*sqlx.DB)

	var data interface{}
	log.Println(ctx.Request().URL)

	err := ctx.Bind(data)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	res, err := connection.NamedExec("INSERT INTO person (id, name, email, address) VALUES (:id, :name, :email, address)", data) // shaper of params depending on resourceType
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusAccepted, res)
}

func HandleRead(ctx echo.Context) {

}

func HandleDelete(ctx echo.Context) {

}

func HandleUpdate(ctx echo.Context) {

}
