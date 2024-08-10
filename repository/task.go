package repository

import (
	"todoList/db"
	"todoList/models"
)

func AddTask(task models.Task) error {
	result := db.GetDBConn().Create(&task)
	return result.Error
}

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	result := db.GetDBConn().Find(&tasks)
	return tasks, result.Error
}

func UpdateTask(taskID uint, title, description string) error {
	var task models.Task
	result := db.GetDBConn().Model(&task).Where(
		"id = ?", taskID).Updates(models.Task{Title: title, Description: description})
	return result.Error
}

func ToggleStatus(taskID uint) error {
	var task models.Task
	result := db.GetDBConn().Model(&task).Where(
		"id = ?", taskID).Update("IsDone", !task.IsDone)
	return result.Error
}

func DeleteTask(taskID uint) error {
	result := db.GetDBConn().Delete(&models.Task{}, taskID)
	return result.Error
}

func InsertExistingData(tasks []models.Task) error {
	for _, task := range tasks {
		result := db.GetDBConn().Create(&task)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func SetPriority(taskID uint, priority int) error {
	var task models.Task
	result := db.GetDBConn().Model(&task).Where(
		"id = ?", taskID).Update("Priority", priority)
	return result.Error
}

func GetTasksByIsDone(status string) ([]models.Task, error) {
	var tasks []models.Task
	result := db.GetDBConn().Where("IsDone = ?", status == "completed").Find(&tasks)
	return tasks, result.Error
}

func SortTasksByDate() ([]models.Task, error) {
	var tasks []models.Task
	result := db.GetDBConn().Order("CreatedAt").Find(&tasks)
	return tasks, result.Error
}

func SortTasksByStatus() ([]models.Task, error) {
	var tasks []models.Task
	result := db.GetDBConn().Order("IsDone").Find(&tasks)
	return tasks, result.Error
}

func SortTasksByPriority() ([]models.Task, error) {
	var tasks []models.Task
	result := db.GetDBConn().Order("Priority").Find(&tasks)
	return tasks, result.Error
}
