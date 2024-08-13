package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var dbConn *sql.DB

func ConnectToDB() error {
	connStr := "user=postgres password=12345678 dbname=todo_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	fmt.Println("Successfully connected to DB")

	dbConn = db
	return nil
}

func CloseDbConnection() error {
	err := dbConn.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDBConn() *sql.DB {
	return dbConn
}
