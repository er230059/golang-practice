package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var dbInstance *sql.DB

func initDb() {
	userName := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", userName, password, host, port, dbName)

	db, err := sql.Open("mysql", connectionString)
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
