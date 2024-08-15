package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RunRoutes() error {
	r := gin.Default()
	r.GET("/ping", PingPong)

	userGroup := r.Group("/users")
	{
		userGroup.GET("/", GetAllUsers)
		userGroup.GET("/:id", GetUserByID)
		userGroup.POST("/", CreateUser)
		userGroup.PUT("/:id", UpdateUser)
		userGroup.DELETE("/:id", DeleteUser)
	}
	taskGroup := r.Group("/tasks")
	{
		taskGroup.GET("/", GetAllTasks)
		taskGroup.GET("/:id", GetTaskByID)
		taskGroup.POST("/", CreateTask)
		taskGroup.PUT("/:id", UpdateTask)
		taskGroup.PATCH("/:id/status", ToggleTaskStatus)
		taskGroup.DELETE("/:id", DeleteTask)
		taskGroup.POST("/insert", InsertExistingTasks)
		taskGroup.PUT("/:id/priority", SetTaskPriority)
		taskGroup.GET("/status/:status", GetTasksByStatus)
		taskGroup.GET("/sort/date", SortTasksByDate)
		taskGroup.GET("/sort/status", SortTasksByStatus)
		taskGroup.GET("/sort/priority", SortTasksByPriority)
	}

	port := ":8181"
	err := r.Run(port)
	if err != nil {
		return err
	}
	return nil
}

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
