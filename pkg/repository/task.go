package repository

import (
	"todoList/db"
	"todoList/models"
)

func GetAllTasks() (tasks []models.Task, err error) {
	err = db.GetDBConn().Where("is_deleted = ?", false).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTaskById(id uint) (task models.Task, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&task).Error
	if err != nil {
		return task, err
	}
	return task, nil
}

func AddTask(task models.Task) error {
	result := db.GetDBConn().Create(&task)
	return result.Error
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

func DeleteTask(id uint) error {
	var task models.Task
	result := db.GetDBConn().Model(&task).Where("id = ?", id).Update("is_deleted", true)
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
	result := db.GetDBConn().Where("is_done = ?", status == "completed").Find(&tasks)
	return tasks, result.Error
}

func SortTasksByDate() ([]models.Task, error) {
	var tasks []models.Task
	result := db.GetDBConn().Order("created_at").Find(&tasks)
	return tasks, result.Error
}

func SortTasksByStatus() ([]models.Task, error) {
	var tasks []models.Task
	result := db.GetDBConn().Order("is_done").Find(&tasks)
	return tasks, result.Error
}

func SortTasksByPriority() ([]models.Task, error) {
	var tasks []models.Task
	result := db.GetDBConn().Order("Priority").Find(&tasks)
	return tasks, result.Error
}
