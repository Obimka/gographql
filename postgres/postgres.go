package postgres

import (
	"cadoles/graphql/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
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
	conf := config.GetConfig()
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.DB_HOST, conf.DB_PORT, conf.DB_USER, conf.DB_PASSWORD, conf.DB_NAME)
	var err error
	DB, err = sql.Open("postgres", dbinfo)
	checkErr(err, DB)
}
func DBClose() {
	DB.Close()
}
