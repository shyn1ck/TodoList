package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"todoList/controllers"
	"todoList/db"
)

func main() {
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

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:63342/TodoList/template/"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))
	router.GET("/tasks", controllers.GetAllTasksHandler)
	router.POST("/tasks", controllers.AddTaskHandler)
	router.PUT("/tasks/:id", controllers.EditTaskHandler)
	router.DELETE("/tasks/:id", controllers.DeleteTaskHandler)
	router.GET("/tasks/filter", controllers.FilterTasksByIsDoneHandler)
	router.GET("/tasks/sort", controllers.SortTasksHandler)
	router.POST("/tasks/test-data", controllers.InsertDataTasksHandler)
	router.PUT("/tasks/:id/toggle", controllers.ToggleTaskStatusHandler)
	router.PUT("/tasks/:id/priority", controllers.SetTaskPriorityHandler)
	port := 8080
	fmt.Printf("Server started on port %d\n", port)
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
