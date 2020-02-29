package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var dbInstance *sql.DB

func initDb() {
	db, err := sql.Open("mysql", "root:er230059@/go")
	if err != nil {
		fmt.Printf("ERR: %s\n", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	dbInstance = db
}

func GetDb() *sql.DB {
	if dbInstance == nil {
		initDb()
	}
	return dbInstance
}
