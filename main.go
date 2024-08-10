package main

import (
	"log"
	"net/http"
	"strings"
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

	// Задаем маршруты для API
	http.HandleFunc("/tasks", TaskHandler)           //  Обработка  всех  задач
	http.HandleFunc("/tasks/", TaskHandler)          //  Обработка  задач  по  ID
	http.HandleFunc("/tasks/toggle", TaskHandler)    //  Переключение  статуса
	http.HandleFunc("/tasks/priority", TaskHandler)  //  Изменение  приоритета
	http.HandleFunc("/tasks/test-data", TaskHandler) //  Вставка  тестовых  данных
	http.HandleFunc("/tasks/filter", TaskHandler)    //  Фильтрация  по  статусу
	http.HandleFunc("/tasks/sort", TaskHandler)      //  Сортировка

	log.Println("Starting server on :63342")
	if err := http.ListenAndServe(":63342", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch {
		case strings.HasPrefix(r.URL.Path, "/tasks/filter"):
			controllers.FilterTasksByIsDoneHandler(w, r)
		case strings.HasPrefix(r.URL.Path, "/tasks/sort"):
			controllers.SortTasksHandler(w, r)
		default:
			controllers.GetAllTasksHandler(w, r)
		}
	case http.MethodPost:
		switch {
		case strings.HasPrefix(r.URL.Path, "/tasks/test-data"):
			controllers.InsertDataTasksHandler(w, r)
		default:
			controllers.AddTaskHandler(w, r)
		}
	case http.MethodPut:
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
