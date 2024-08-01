package repository

import (
	"fmt"
	"todoList/db"
	"todoList/models"
)

func AddTask(task models.Task) error {
	_, err := db.GetDBConn().Exec(db.AddNewTaskQuery, task.Title, task.Description, task.Priority)
	if err != nil {
		return err
	}
	return nil
}

func GetAllTasks() (tasks []models.Task, err error) {
	rows, err := db.GetDBConn().Query(db.GetAllTasksQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.IsDone, &task.Priority, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func UpdateTask(taskID int, title, description string) error {
	_, err := db.GetDBConn().Exec(db.UpdateTaskQuery, title, description, taskID)
	return err
}

func ToggleStatus(taskID int) error {
	_, err := db.GetDBConn().Exec(db.ToggleStatusQuery, taskID)
	return err
}

func DeleteTask(taskID int) error {
	_, err := db.GetDBConn().Exec(db.DeleteTaskQuery, taskID)
	return err
}

func SetPriority(taskID int, priority int) error {
	_, err := db.GetDBConn().Exec(db.SetPriorityQuery, priority, taskID)
	return err
}

func GetTasksByIsDone(status string) (tasks []models.Task, err error) {
	var isDone bool
	if status == "completed" {
		isDone = true
	} else if status == "not completed" {
		isDone = false
	} else {
		return nil, fmt.Errorf("invalid status: %s", status)
	}

	rows, err := db.GetDBConn().Query(db.GetTasksByIsDoneQuery, isDone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.IsDone, &task.Priority, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
func InsertExistingData() error {
	_, err := db.GetDBConn().Exec(db.InsertExistingData)
	if err != nil {
		return err
	}
	return nil
}

func SortTasksByDate() ([]models.Task, error) {
	return sortTasks(db.SortByDateQuery)
}

func SortTasksByStatus() ([]models.Task, error) {
	return sortTasks(db.SortByStatusQuery)
}

func SortTasksByPriority() ([]models.Task, error) {
	return sortTasks(db.SortByPriorityQuery)
}

func sortTasks(query string) ([]models.Task, error) {
	rows, err := db.GetDBConn().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.IsDone, &task.Priority, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
