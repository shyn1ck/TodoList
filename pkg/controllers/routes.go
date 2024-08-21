package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"todoList/configs"
)

func RunRoutes() error {
	r := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)

	r.GET("/ping", PingPong)

	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}

	userGroup := r.Group("/users")
	userGroup.Use(checkUserAuthentication)
	{
		userGroup.GET("/", GetAllUsers)
		userGroup.GET("/:id", GetUserByID)
		userGroup.POST("/", CreateUser)
		userGroup.PUT("/:id", UpdateUser)
		userGroup.DELETE("/:id", DeleteUser)
	}

	taskGroup := r.Group("/tasks")
	taskGroup.Use(checkUserAuthentication)
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

	err := r.Run(fmt.Sprintf("%s:%s", configs.AppSettings.AppParams.ServerURL, configs.AppSettings.AppParams.PortRun))

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
