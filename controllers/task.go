package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todoList/models"
	"todoList/repository"
)

func AddTask() {
	var task models.Task
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите заголовок задачи:")
	title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при чтении заголовка:", err)
		return
	}
	task.Title = strings.TrimSpace(title)
	fmt.Println("Введите описание задачи:")
	description, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при чтении описания:", err)
		return
	}
	task.Description = strings.TrimSpace(description)
	fmt.Println("Введите приоритет задачи (целое число):")
	priorityStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при чтении приоритета:", err)
		return
	}
	priorityStr = strings.TrimSpace(priorityStr)
	task.Priority, err = strconv.Atoi(priorityStr)
	if err != nil {
		fmt.Println("Ошибка при преобразовании приоритета в целое число:", err)
		return
	}
	err = repository.AddTask(task)
	if err != nil {
		fmt.Println("Ошибка при добавлении задачи:", err)
		return
	}

	fmt.Println("Задача успешно добавлена!")
}

func GetAllTasks() {
	tasks, err := repository.GetAllTasks()
	if err != nil {
		fmt.Println("Error retrieving tasks:", err)
		return
	}

	fmt.Println("Tasks:")
	for _, task := range tasks {
		status := "Not Completed"
		if task.IsDone {
			status = "Completed"
		}
		fmt.Printf("ID: %d, Title: %s, Description: %s, Status: %s, Priority: %d, Created At: %s\n",
			task.ID, task.Title, task.Description, status, task.Priority, task.CreatedAt)
	}
}

func EditTask() {
	var taskID int
	var newTitle, newDescription string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите ID задачи для редактирования:")
	_, err := fmt.Scan(&taskID)
	if err != nil {
		fmt.Println("Ошибка при чтении ID задачи:", err)
		return
	}
	fmt.Println("Введите новый заголовок:")
	title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при чтении нового заголовка:", err)
		return
	}
	newTitle = strings.TrimSpace(title)
	fmt.Println("Введите новое описание:")
	description, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при чтении нового описания:", err)
		return
	}
	newDescription = strings.TrimSpace(description)
	err = repository.UpdateTask(taskID, newTitle, newDescription)
	if err != nil {
		fmt.Println("Ошибка при обновлении задачи:", err)
		return
	}

	fmt.Println("Задача успешно обновлена!")
}

func ToggleTaskStatus() {
	var taskID int

	fmt.Println("Enter the task ID to toggle status:")
	_, err := fmt.Scan(&taskID)
	if err != nil {
		fmt.Println("Error reading task ID:", err)
		return
	}

	err = repository.ToggleStatus(taskID)
	if err != nil {
		fmt.Println("Error toggling task status:", err)
		return
	}

	fmt.Println("Task status toggled successfully!")
}

func DeleteTask() {
	var taskID int

	fmt.Println("Enter the task ID to delete:")
	_, err := fmt.Scan(&taskID)
	if err != nil {
		fmt.Println("Error reading task ID:", err)
		return
	}

	err = repository.DeleteTask(taskID)
	if err != nil {
		fmt.Println("Error deleting task:", err)
		return
	}

	fmt.Println("Task deleted successfully!")
}

func InsertDataTasks() {
	err := repository.InsertExistingData()
	if err != nil {
		fmt.Println("Error inserting test tasks:", err)
		return
	}
	fmt.Println("Test tasks inserted successfully!")
}

func SetTaskPriority() {
	var taskID int
	var priority int

	fmt.Println("Enter the task ID to set priority:")
	_, err := fmt.Scan(&taskID)
	if err != nil {
		fmt.Println("Error reading task ID:", err)
		return
	}

	fmt.Println("Enter the priority level (1-5):")
	_, err = fmt.Scan(&priority)
	if err != nil {
		fmt.Println("Error reading priority level:", err)
		return
	}

	err = repository.SetPriority(taskID, priority)
	if err != nil {
		fmt.Println("Error setting task priority:", err)
		return
	}

	fmt.Println("Task priority set successfully!")
}

func FilterTasksByIsDone() {
	var status string

	fmt.Println("Enter the status to filter by (completed/not completed):")
	_, err := fmt.Scan(&status)
	if err != nil {
		fmt.Println("Error reading status:", err)
		return
	}

	status = strings.ToLower(status)
	if status != "completed" && status != "not completed" {
		fmt.Println("Invalid status. Please enter 'completed' or 'not completed'.")
		return
	}

	tasks, err := repository.GetTasksByIsDone(status)
	if err != nil {
		fmt.Println("Error filtering tasks:", err)
		return
	}

	fmt.Println("Filtered Tasks:")
	for _, task := range tasks {
		status := "Not Completed"
		if task.IsDone {
			status = "Completed"
		}
		fmt.Printf("ID: %d, Title: %s, Description: %s, Status: %s, Priority: %d, Created At: %s\n",
			task.ID, task.Title, task.Description, status, task.Priority, task.CreatedAt)
	}
}

func SortTasks() {
	var sortOption string

	fmt.Println("Введите опцию сортировки (date/status/priority):")
	_, err := fmt.Scan(&sortOption)
	if err != nil {
		fmt.Println("Ошибка при чтении опции сортировки:", err)
		return
	}

	var tasks []models.Task
	switch strings.ToLower(sortOption) {
	case "date":
		tasks, err = repository.SortTasksByDate()
	case "status":
		tasks, err = repository.SortTasksByStatus()
	case "priority":
		tasks, err = repository.SortTasksByPriority()
	default:
		fmt.Println("Неверная опция сортировки. Используйте 'date', 'status' или 'priority'.")
		return
	}

	if err != nil {
		fmt.Println("Ошибка при сортировке задач:", err)
		return
	}

	fmt.Println("Отсортированные задачи:")
	for _, task := range tasks {
		status := "Не выполнена"
		if task.IsDone {
			status = "Выполнена"
		}
		fmt.Printf("ID: %d, Заголовок: %s, Описание: %s, Статус: %s, Приоритет: %d, Создано: %s\n",
			task.ID, task.Title, task.Description, status, task.Priority, task.CreatedAt)
	}
}
