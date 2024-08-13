package db

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func ConnectToDB() error {
	connStr := "user=postgres password=12345678 dbname=todo_db sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("Connected to database")

	dbConn = db
	return nil
}

func CloseDbConnection() error {
	sqlDB, err := dbConn.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDBConn() *gorm.DB {
	return dbConn
}
