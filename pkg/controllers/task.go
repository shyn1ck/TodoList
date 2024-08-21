package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todoList/logger"
	"todoList/models"
	"todoList/pkg/service"
)

func CreateTask(c *gin.Context) {
	logger.Info.Println("[CreateTask]: Received request to create a task from IP:", c.ClientIP())
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	err := service.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	logger.Info.Println("[CreateTask]: Task created successfully by IP:", c.ClientIP())
	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully"})
}

func GetAllTasks(c *gin.Context) {
	logger.Info.Println("[GetAllTasks]: Received request to get all tasks from IP:", c.ClientIP())
	tasks, err := service.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	logger.Info.Println("[GetAllTasks]: Tasks retrieved successfully by IP:", c.ClientIP())
	c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	logger.Info.Println("[GetTaskByID]: Received request to get task by ID from IP:", c.ClientIP())
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task ID"})
		return
	}

	task, err := service.GetTaskByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found"})
		return
	}

	logger.Info.Printf("[GetTaskByID]: Task retrieved successfully, ID: %v, IP: %v\n", id, c.ClientIP())
	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	logger.Info.Println("[UpdateTask]: Received request to update task from IP:", c.ClientIP())
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task ID"})
		return
	}

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	err = service.UpdateTask(uint(id), input.Title, input.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	logger.Info.Printf("[UpdateTask]: Task updated successfully, ID: %v, IP: %v\n", id, c.ClientIP())
	c.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully"})
}

func ToggleTaskStatus(c *gin.Context) {
	logger.Info.Println("[ToggleTaskStatus]: Received request to toggle task status from IP:", c.ClientIP())
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task ID"})
		return
	}

	err = service.ToggleTaskStatus(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	logger.Info.Printf("[ToggleTaskStatus]: Task status toggled successfully, ID: %v, IP: %v\n", id, c.ClientIP())
	c.JSON(http.StatusOK, gin.H{
		"message": "Task status toggled successfully"})
}

func DeleteTask(c *gin.Context) {
	logger.Info.Println("[DeleteTask]: Received request to delete task from IP:", c.ClientIP())
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task ID"})
		return
	}

	err = service.DeleteTask(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	logger.Info.Printf("[DeleteTask]: Task deleted successfully, ID: %v, IP: %v\n", id, c.ClientIP())
	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully"})
}

func InsertExistingTasks(c *gin.Context) {
	logger.Info.Println("[InsertExistingTasks]: Received request to insert existing tasks from IP:", c.ClientIP())
	var tasks []models.Task
	if err := c.ShouldBindJSON(&tasks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	err := service.InsertExistingTasks(tasks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	logger.Info.Println("[InsertExistingTasks]: Tasks inserted successfully by IP:", c.ClientIP())
	c.JSON(http.StatusCreated, gin.H{
		"message": "Tasks inserted successfully"})
}

func SetTaskPriority(c *gin.Context) {
	logger.Info.Println("[SetTaskPriority]: Received request to set task priority from IP:", c.ClientIP())
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task ID"})
		return
	}

	var input struct {
		Priority int `json:"priority"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	err = service.SetTaskPriority(uint(id), input.Priority)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	logger.Info.Printf("[SetTaskPriority]: Task priority set successfully, ID: %v, IP: %v\n", id, c.ClientIP())
	c.JSON(http.StatusOK, gin.H{
		"message": "Task priority set successfully"})
}

func GetTasksByStatus(c *gin.Context) {
	logger.Info.Println("[GetTasksByStatus]: Received request to get tasks by status from IP:", c.ClientIP())
	status := c.Param("status")
	tasks, err := service.GetTasksByStatus(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	logger.Info.Printf("[GetTasksByStatus]: Tasks retrieved successfully by status, status: %v, IP: %v\n", status, c.ClientIP())
	c.JSON(http.StatusOK, tasks)
}

func SortTasksByDate(c *gin.Context) {
	logger.Info.Println("[SortTasksByDate]: Received request to sort tasks by date from IP:", c.ClientIP())
	tasks, err := service.SortTasksByDate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	logger.Info.Println("[SortTasksByDate]: Tasks sorted by date successfully by IP:", c.ClientIP())
	c.JSON(http.StatusOK, tasks)
}

func SortTasksByStatus(c *gin.Context) {
	logger.Info.Println("[SortTasksByStatus]: Received request to sort tasks by status from IP:", c.ClientIP())
	tasks, err := service.SortTasksByStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	logger.Info.Println("[SortTasksByStatus]: Tasks sorted by status successfully by IP:", c.ClientIP())
	c.JSON(http.StatusOK, tasks)
}

func SortTasksByPriority(c *gin.Context) {
	logger.Info.Println("[SortTasksByPriority]: Received request to sort tasks by priority from IP:", c.ClientIP())
	tasks, err := service.SortTasksByPriority()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	logger.Info.Println("[SortTasksByPriority]: Tasks sorted by priority successfully by IP:", c.ClientIP())
	c.JSON(http.StatusOK, tasks)
}
