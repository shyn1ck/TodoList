package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var dbConn *sql.DB

func ConnectToDB() error {
	connStr := "user=postgres password=2003 dbname=todo_list_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	fmt.Println("Connected to database")

	dbConn = db
	return nil
}

func CloseDBConn() error {
	err := dbConn.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDBConn() *sql.DB {
	return dbConn
}
