package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todoList/models"
	"todoList/pkg/service"
)

func CreateTask(c *gin.Context) {
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

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully"})
}

func GetAllTasks(c *gin.Context) {
	tasks, err := service.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
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

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully"})
}

func ToggleTaskStatus(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Task status toggled successfully"})
}

func DeleteTask(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully"})
}

func InsertExistingTasks(c *gin.Context) {
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

	c.JSON(http.StatusCreated, gin.H{
		"message": "Tasks inserted successfully"})
}

func SetTaskPriority(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Task priority set successfully"})
}

func GetTasksByStatus(c *gin.Context) {
	status := c.Param("status")
	tasks, err := service.GetTasksByStatus(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func SortTasksByDate(c *gin.Context) {
	tasks, err := service.SortTasksByDate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func SortTasksByStatus(c *gin.Context) {
	tasks, err := service.SortTasksByStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func SortTasksByPriority(c *gin.Context) {
	tasks, err := service.SortTasksByPriority()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
