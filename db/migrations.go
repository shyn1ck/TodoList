package db

import (
	"errors"
	"log"
	"todoList/models"
)

func Migrate() error {
	if dbConn == nil {
		return errors.New("database connection is not initialized")
	}

	err := dbConn.AutoMigrate(&models.Task{}, &models.User{})
	if err != nil {
		return errors.New("failed to migrate database schema: " + err.Error())
	}

	log.Println("Database migration completed successfully")
	return nil
}
