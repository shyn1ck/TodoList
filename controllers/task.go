// controllers/tasks.go

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"todoList/models"
	"todoList/repository"
)

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var task models.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, `{"error":"Invalid request body"}`, http.StatusBadRequest)
			return
		}

		if err := repository.AddTask(task); err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
}

func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := repository.GetAllTasks()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}
}

func EditTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		var task models.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, `{"error":"Invalid request body"}`, http.StatusBadRequest)
			return
		}

		taskIDStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
		taskID, err := strconv.Atoi(taskIDStr)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, `{"error":"Invalid task ID"}`, http.StatusBadRequest)
			return
		}

		task.ID = uint(taskID)

		if err := repository.UpdateTask(task.ID, task.Title, task.Description); err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
}

func ToggleTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	taskIDStr := r.URL.Query().Get("id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error":"Invalid task ID"}`, http.StatusBadRequest)
		return
	}

	if err := repository.ToggleStatus(uint(taskID)); err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskIDStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error":"Invalid task ID"}`, http.StatusBadRequest)
		return
	}

	if err := repository.DeleteTask(uint(taskID)); err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func InsertDataTasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		tasks := []models.Task{
			{Title: "Изучить основы Go", Description: "Пройти курс по основам Go на платформе онлайн-обучения", Priority: 1, IsDone: true},
			{Title: "Сделать домашнее задание по алгоритмам", Description: "Решить задачи по алгоритмам и структурам данных", Priority: 2, IsDone: true},
			{Title: "Работа с базами данных", Description: "Изучить SQL и ORM для работы с базами данных", Priority: 3},
			{Title: "Разработка API", Description: "Разработать REST API для учебного проекта", Priority: 2, IsDone: true},
			{Title: "Прочитать книгу по программированию", Description: "Прочитать книгу 'Чистый код' Роберта Мартина", Priority: 1},
			{Title: "Участие в хакатоне", Description: "Принять участие в студенческом хакатоне", Priority: 3},
			{Title: "Создать GitHub репозиторий", Description: "Создать и настроить репозиторий для учебных проектов на GitHub", Priority: 2},
			{Title: "Установить среду разработки", Description: "Настроить IDE и все необходимые плагины для работы", Priority: 1, IsDone: true},
			{Title: "Проектирование системы", Description: "Разработать архитектуру системы для учебного проекта", Priority: 3, IsDone: true},
			{Title: "Учебный проект: веб-приложение", Description: "Создать простое веб-приложение с использованием фреймворка", Priority: 2},
		}

		if err := repository.InsertExistingData(tasks); err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"Test tasks successfully inserted!"}`))
	}
}

func SetTaskPriorityHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {
		var data struct {
			ID       uint `json:"id"`
			Priority int  `json:"priority"`
		}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, `{"error":"Invalid request body"}`, http.StatusBadRequest)
			return
		}

		if err := repository.SetPriority(data.ID, data.Priority); err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
}

func FilterTasksByIsDoneHandler(w http.ResponseWriter, r *http.Request) {
	status := strings.ToLower(r.URL.Query().Get("status"))

	if status != "completed" && status != "not completed" {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error":"Invalid status"}`, http.StatusBadRequest)
		return
	}

	tasks, err := repository.GetTasksByIsDone(status)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}
}

func SortTasksHandler(w http.ResponseWriter, r *http.Request) {
	sortOption := strings.ToLower(r.URL.Query().Get("option"))

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
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error":"Invalid sort option"}`, http.StatusBadRequest)
		return
	}

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}
}
