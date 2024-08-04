package repository

import (
	"todoList/db"
	"todoList/models"
)

func AddTask(task models.Task) error {
	return db.GetDBConn().Create(&task).Error
}

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := db.GetDBConn().Find(&tasks).Error
	return tasks, err
}

func UpdateTask(taskID uint, title, description string) error {
	return db.GetDBConn().Model(&models.Task{}).Where("id = ?", taskID).Updates(models.Task{Title: title, Description: description}).Error
}

func ToggleStatus(taskID uint) error {
	var task models.Task
	err := db.GetDBConn().First(&task, taskID).Error
	if err != nil {
		return err
	}
	task.IsDone = !task.IsDone
	return db.GetDBConn().Save(&task).Error
}

func DeleteTask(taskID uint) error {
	return db.GetDBConn().Delete(&models.Task{}, taskID).Error
}

func SetPriority(taskID uint, priority int) error {
	return db.GetDBConn().Model(&models.Task{}).Where("id = ?", taskID).Update("priority", priority).Error
}

func GetTasksByIsDone(status string) ([]models.Task, error) {
	var tasks []models.Task
	isDone := status == "completed"
	err := db.GetDBConn().Where("is_done = ?", isDone).Find(&tasks).Error
	return tasks, err
}

func InsertExistingData(tasks []models.Task) error {
	return db.GetDBConn().Create(&tasks).Error
}

func SortTasksByDate() ([]models.Task, error) {
	return sortTasks("created_at")
}

func SortTasksByStatus() ([]models.Task, error) {
	return sortTasks("is_done")
}

func SortTasksByPriority() ([]models.Task, error) {
	return sortTasks("priority")
}

func sortTasks(orderBy string) ([]models.Task, error) {
	var tasks []models.Task
	err := db.GetDBConn().Order(orderBy).Find(&tasks).Error
	return tasks, err
}
