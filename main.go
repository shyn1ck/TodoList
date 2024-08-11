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

	http.HandleFunc("/tasks", TaskHandler)
	http.HandleFunc("/tasks/", TaskHandler)
	http.HandleFunc("/tasks/toggle", TaskHandler)
	http.HandleFunc("/tasks/priority", TaskHandler)
	http.HandleFunc("/tasks/test-data", TaskHandler)
	http.HandleFunc("/tasks/filter", TaskHandler)
	http.HandleFunc("/tasks/sort", TaskHandler)

	log.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		switch {
		case strings.HasPrefix(r.URL.Path, "/tasks/filter"):
			controllers.FilterTasksByIsDoneHandler(w, r)
		case strings.HasPrefix(r.URL.Path, "/tasks/sort"):
			controllers.SortTasksHandler(w, r)
		default:
			controllers.GetAllTasksHandler(w, r)
			w.Write([]byte(`{"message":"Tasks retrieved successfully"}`))
		}
	case http.MethodPost:
		switch {
		case strings.HasPrefix(r.URL.Path, "/tasks/test-data"):
			controllers.InsertDataTasksHandler(w, r)
			w.Write([]byte(`{"message":"Test data inserted successfully"}`))
		default:
			controllers.AddTaskHandler(w, r)
			w.Write([]byte(`{"message":"Task added successfully"}`))
		}
	case http.MethodPut:
		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) > 2 && pathParts[1] == "tasks" {
			if pathParts[2] == "toggle" {
				controllers.ToggleTaskStatusHandler(w, r)
				w.Write([]byte(`{"message":"Task status toggled successfully"}`))
			} else if pathParts[2] == "priority" {
				controllers.SetTaskPriorityHandler(w, r)
				w.Write([]byte(`{"message":"Task priority set successfully"}`))
			} else {
				controllers.EditTaskHandler(w, r)
				w.Write([]byte(`{"message":"Task edited successfully"}`))
			}
		} else {
			http.Error(w, `{"error":"Not Found"}`, http.StatusNotFound)
		}
	case http.MethodDelete:
		if strings.HasPrefix(r.URL.Path, "/tasks/") && !strings.Contains(r.URL.Path, "/toggle") && !strings.Contains(r.URL.Path, "/priority") {
			controllers.DeleteTaskHandler(w, r)
			w.Write([]byte(`{"message":"Task deleted successfully"}`))
		} else {
			http.Error(w, `{"error":"Not Found"}`, http.StatusNotFound)
		}
	default:
		http.Error(w, `{"error":"Unsupported method"}`, http.StatusMethodNotAllowed)
	}
}
