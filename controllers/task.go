package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"todoList/models"
	"todoList/repository"
)

func AddTaskHandler(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный запрос"})
		return
	}

	if err := repository.AddTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Задача успешно добавлена"})
}

func GetAllTasksHandler(c *gin.Context) {
	tasks, err := repository.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func EditTaskHandler(c *gin.Context) {
	taskIDStr := c.Param("id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID задачи"})
		return
	}

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный запрос"})
		return
	}

	task.ID = uint(taskID)

	if err := repository.UpdateTask(task.ID, task.Title, task.Description); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Задача успешно отредактирована"})
}

func ToggleTaskStatusHandler(c *gin.Context) {
	taskIDStr := c.Param("id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID задачи"})
		return
	}

	if err := repository.ToggleStatus(uint(taskID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Статус задачи успешно изменен"})
}

func DeleteTaskHandler(c *gin.Context) {
	taskIDStr := c.Param("id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID задачи"})
		return
	}

	if err := repository.DeleteTask(uint(taskID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Задача успешно удалена"})
}

func InsertDataTasksHandler(c *gin.Context) {
	tasks := []models.Task{
		{Title: "Изучить основы Go", Description: "Пройти курс по основам Go на онлайн-платформе", Priority: 1},
		{Title: "Сделать домашнюю работу по алгоритмам", Description: "Решить задачи по алгоритмам и структурам данных", Priority: 2},
		{Title: "Изучить основы Go", Description: "Пройти курс по основам Go на онлайн-платформе", Priority: 1, IsDone: true},
		{Title: "Сделать домашнюю работу по алгоритмам", Description: "Решить задачи по алгоритмам и структурам данных", Priority: 2, IsDone: true},
		{Title: "Работа с базами данных", Description: "Изучить SQL и ORM для работы с базами данных", Priority: 3},
		{Title: "Разработка API", Description: "Разработать REST API для учебного проекта", Priority: 2},
		{Title: "Разработка API", Description: "Разработать REST API для учебного проекта", Priority: 2, IsDone: true},
		{Title: "Прочитать книгу по программированию", Description: "Прочитать книгу 'Чистый код' Роберта Мартина", Priority: 1},
		{Title: "Участие в хакатоне", Description: "Принять участие в студенческом хакатоне", Priority: 3},
		{Title: "Создать репозиторий на GitHub", Description: "Создать и настроить репозиторий для учебных проектов на GitHub", Priority: 2},
		{Title: "Настроить среду разработки", Description: "Настроить IDE и необходимые плагины для работы", Priority: 1},
		{Title: "Системный дизайн", Description: "Разработать системную архитектуру для учебного проекта", Priority: 3},
		{Title: "Настроить среду разработки", Description: "Настроить IDE и необходимые плагины для работы", Priority: 1, IsDone: true},
		{Title: "Системный дизайн", Description: "Разработать системную архитектуру для учебного проекта", Priority: 3, IsDone: true},
		{Title: "Учебный проект: Веб-приложение", Description: "Создать простое веб-приложение с использованием фреймворка", Priority: 2},
	}

	if err := repository.InsertExistingData(tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Тестовые задачи успешно добавлены!"})
}

func SetTaskPriorityHandler(c *gin.Context) {
	var data struct {
		ID       uint `json:"id"`
		Priority int  `json:"priority"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный запрос"})
		return
	}

	if err := repository.SetPriority(data.ID, data.Priority); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Приоритет задачи успешно изменен"})
}

func FilterTasksByIsDoneHandler(c *gin.Context) {
	status := strings.ToLower(c.Query("status"))

	if status != "completed" && status != "not completed" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный статус"})
		return
	}

	tasks, err := repository.GetTasksByIsDone(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func SortTasksHandler(c *gin.Context) {
	sortOption := strings.ToLower(c.Query("option"))

	var tasks []models.Task
	var err error

	switch sortOption {
	case "date":
		tasks, err = repository.SortTasksByDate()
	case "status":
		tasks, err = repository.SortTasksByStatus()
	case "priority":
		tasks, err = repository.SortTasksByPriority()
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный параметр сортировки"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
