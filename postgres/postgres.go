package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "graphql"
	DB_PASSWORD = "graphql"
	DB_NAME     = "graphql"
)

var (
	DB *sql.DB
)

func checkErr(err error, DB *sql.DB) {
	if err != nil {
		panic(err)
	}
}

func DBConnect() {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	var err error
	DB, err = sql.Open("postgres", dbinfo)
	checkErr(err, DB)
}
func DBClose() {
	DB.Close()
}
