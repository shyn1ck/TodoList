package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"todoList/models"
	"todoList/pkg/service"
)

func CreateTask(c *gin.Context) {
	log.Println("controllers.CreateTask: Received request to create a task")
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		log.Printf("controllers.CreateTask: Failed to bind JSON, error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	err := service.CreateTask(task)
	if err != nil {
		log.Printf("controllers.CreateTask: Failed to create task, error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	log.Println("controllers.CreateTask: Task created successfully")
	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully"})
}

func GetAllTasks(c *gin.Context) {
	log.Println("controllers.GetAllTasks: Received request to get all tasks")
	tasks, err := service.GetAllTasks()
	if err != nil {
		log.Printf("controllers.GetAllTasks: Failed to retrieve tasks, error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	log.Println("controllers.GetAllTasks: Tasks retrieved successfully")
	c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	log.Println("controllers.GetTaskByID: Received request to get task by ID")
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Printf("controllers.GetTaskByID: Invalid task ID, error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task ID"})
		return
	}

	task, err := service.GetTaskByID(uint(id))
	if err != nil {
		log.Printf("controllers.GetTaskByID: Task not found, ID: %v, error: %v\n", id, err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found"})
		return
	}

	log.Printf("controllers.GetTaskByID: Task retrieved successfully, ID: %v\n", id)
	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	log.Println("controllers.UpdateTask: Received request to update task")
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Printf("controllers.UpdateTask: Invalid task ID, error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task ID"})
		return
	}

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("controllers.UpdateTask: Failed to bind JSON, error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	err = service.UpdateTask(uint(id), input.Title, input.Description)
	if err != nil {
		log.Printf("controllers.UpdateTask: Failed to update task, ID: %v, error: %v\n", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	log.Printf("controllers.UpdateTask: Task updated successfully, ID: %v\n", id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully"})
}

func ToggleTaskStatus(c *gin.Context) {
	log.Println("controllers.ToggleTaskStatus: Received request to toggle task status")
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Printf("controllers.ToggleTaskStatus: Invalid task ID, error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task ID"})
		return
	}

	err = service.ToggleTaskStatus(uint(id))
	if err != nil {
		log.Printf("controllers.ToggleTaskStatus: Failed to toggle task status, ID: %v, error: %v\n", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	log.Printf("controllers.ToggleTaskStatus: Task status toggled successfully, ID: %v\n", id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Task status toggled successfully"})
}

func DeleteTask(c *gin.Context) {
	log.Println("controllers.DeleteTask: Received request to delete task")
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Printf("controllers.DeleteTask: Invalid task ID, error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task ID"})
		return
	}

	err = service.DeleteTask(uint(id))
	if err != nil {
		log.Printf("controllers.DeleteTask: Failed to delete task, ID: %v, error: %v\n", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	log.Printf("controllers.DeleteTask: Task deleted successfully, ID: %v\n", id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully"})
}

func InsertExistingTasks(c *gin.Context) {
	log.Println("controllers.InsertExistingTasks: Received request to insert existing tasks")
	var tasks []models.Task
	if err := c.ShouldBindJSON(&tasks); err != nil {
		log.Printf("controllers.InsertExistingTasks: Failed to bind JSON, error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	err := service.InsertExistingTasks(tasks)
	if err != nil {
		log.Printf("controllers.InsertExistingTasks: Failed to insert tasks, error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	log.Println("controllers.InsertExistingTasks: Tasks inserted successfully")
	c.JSON(http.StatusCreated, gin.H{
		"message": "Tasks inserted successfully"})
}

func SetTaskPriority(c *gin.Context) {
	log.Println("controllers.SetTaskPriority: Received request to set task priority")
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Printf("controllers.SetTaskPriority: Invalid task ID, error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid task ID"})
		return
	}

	var input struct {
		Priority int `json:"priority"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("controllers.SetTaskPriority: Failed to bind JSON, error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	err = service.SetTaskPriority(uint(id), input.Priority)
	if err != nil {
		log.Printf("controllers.SetTaskPriority: Failed to set task priority, ID: %v, error: %v\n", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	log.Printf("controllers.SetTaskPriority: Task priority set successfully, ID: %v\n", id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Task priority set successfully"})
}

func GetTasksByStatus(c *gin.Context) {
	log.Println("controllers.GetTasksByStatus: Received request to get tasks by status")
	status := c.Param("status")
	tasks, err := service.GetTasksByStatus(status)
	if err != nil {
		log.Printf("controllers.GetTasksByStatus: Failed to retrieve tasks by status, status: %v, error: %v\n", status, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	log.Printf("controllers.GetTasksByStatus: Tasks retrieved successfully by status, status: %v\n", status)
	c.JSON(http.StatusOK, tasks)
}

func SortTasksByDate(c *gin.Context) {
	log.Println("controllers.SortTasksByDate: Received request to sort tasks by date")
	tasks, err := service.SortTasksByDate()
	if err != nil {
		log.Printf("controllers.SortTasksByDate: Failed to sort tasks by date, error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	log.Println("controllers.SortTasksByDate: Tasks sorted by date successfully")
	c.JSON(http.StatusOK, tasks)
}

func SortTasksByStatus(c *gin.Context) {
	log.Println("controllers.SortTasksByStatus: Received request to sort tasks by status")
	tasks, err := service.SortTasksByStatus()
	if err != nil {
		log.Printf("controllers.SortTasksByStatus: Failed to sort tasks by status, error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	log.Println("controllers.SortTasksByStatus: Tasks sorted by status successfully")
	c.JSON(http.StatusOK, tasks)
}

func SortTasksByPriority(c *gin.Context) {
	log.Println("controllers.SortTasksByPriority: Received request to sort tasks by priority")
	tasks, err := service.SortTasksByPriority()
	if err != nil {
		log.Printf("controllers.SortTasksByPriority: Failed to sort tasks by priority, error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	log.Println("controllers.SortTasksByPriority: Tasks sorted by priority successfully")
	c.JSON(http.StatusOK, tasks)
}
