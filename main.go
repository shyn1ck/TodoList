package main

import (
	"log"
	"net/http"
	"strings"
	"todoList/controllers"
	"todoList/db"
)

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

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

	// Задаем маршруты для API
	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/api/tasks", TaskHandler)           //  Обработка  всех  задач
	http.HandleFunc("/api/tasks/", TaskHandler)          //  Обработка  задач  по  ID
	http.HandleFunc("/api/tasks/toggle", TaskHandler)    //  Переключение  статуса
	http.HandleFunc("/api/tasks/priority", TaskHandler)  //  Изменение  приоритета
	http.HandleFunc("/api/tasks/test-data", TaskHandler) //  Вставка  тестовых  данных
	http.HandleFunc("/api/tasks/filter", TaskHandler)    //  Фильтрация  по  статусу
	http.HandleFunc("/api/tasks/sort", TaskHandler)      //  Сортировка

	log.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		setHeaders(w)
		switch {
		case strings.HasPrefix(r.URL.Path, "/tasks/filter"):
			controllers.FilterTasksByIsDoneHandler(w, r)
		case strings.HasPrefix(r.URL.Path, "/tasks/sort"):
			controllers.SortTasksHandler(w, r)
		default:
			controllers.GetAllTasksHandler(w, r)
		}
	case http.MethodPost:
		setHeaders(w)
		switch {
		case strings.HasPrefix(r.URL.Path, "/tasks/test-data"):
			controllers.InsertDataTasksHandler(w, r)
		default:
			controllers.AddTaskHandler(w, r)
		}
	case http.MethodPut:
		setHeaders(w)
		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) > 2 && pathParts[1] == "tasks" {
			if pathParts[2] == "toggle" {
				controllers.ToggleTaskStatusHandler(w, r)
			} else if pathParts[2] == "priority" {
				controllers.SetTaskPriorityHandler(w, r)
			} else {
				controllers.EditTaskHandler(w, r)
			}
		} else {
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	case http.MethodDelete:
		setHeaders(w)
		if strings.HasPrefix(r.URL.Path, "/tasks/") && !strings.Contains(r.URL.Path,
			"/toggle") && !strings.Contains(r.URL.Path,
			"/priority") {
			controllers.DeleteTaskHandler(w, r)
		} else {
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}
