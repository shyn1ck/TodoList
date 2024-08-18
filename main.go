package main

import (
	"fmt"
	"log"
	"todoList/db"
	"todoList/logger"
	"todoList/pkg/controllers"
)

func main() {
	err := logger.Init()
	if err != nil {
		return
	}
	if err := db.ConnectToDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		if err := db.CloseDBConn(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	if err := db.Migrate(); err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	fmt.Printf("Server started on port %d\n", 8181)
	err = controllers.RunRoutes()
	if err != nil {
		log.Fatal(err)
	}
}
