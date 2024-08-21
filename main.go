package main

import (
	"todoList/db"
	"todoList/logger"
	"todoList/pkg/controllers"
)

func main() {
	err := logger.Init()
	if err != nil {
		logger.Error.Fatalf("Failed to initialize logger: %v", err)
		return
	}
	logger.Info.Println("Logger initialized successfully")
	if err := db.ConnectToDB(); err != nil {
		logger.Error.Fatalf("Failed to connect to database: %v", err)
	}
	logger.Info.Println("Connected to the database successfully")

	defer func() {
		if err := db.CloseDBConn(); err != nil {
			logger.Error.Printf("Error closing database connection: %v", err)
		} else {
			logger.Info.Println("Database connection closed successfully")
		}
	}()
	if err := db.Migrate(); err != nil {
		logger.Error.Fatalf("Failed to run database migrations: %v", err)
	}
	logger.Info.Println("Database migrations ran successfully")
	logger.Info.Printf("Server started on port %d", 8181)
	err = controllers.RunRoutes()
	if err != nil {
		logger.Error.Fatal(err)
	}
}
