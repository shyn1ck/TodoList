package db

import (
	"errors"
	"log"
)

func Migrate() error {
	if dbConn == nil {
		return errors.New("Database connection is not initialized")
	}

	tx, err := dbConn.Begin()
	if err != nil {
		return errors.New("Failed to begin transaction: " + err.Error())
	}
	_, err = tx.Exec(CreateTasksTableQuery)
	if err != nil {
		tx.Rollback()
		return errors.New("Failed to create tasks table: " + err.Error())
	}
	err = tx.Commit()
	if err != nil {
		return errors.New("Failed to commit transaction: " + err.Error())
	}
	log.Println("Database migration completed successfully")
	return nil
}
