package repository

import (
	"todoList/db"
	"todoList/logger"
	"todoList/models"
)

func GetAllTasks() (tasks []models.Task, err error) {
	err = db.GetDBConn().Where("is_deleted = ?", false).Find(&tasks).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllTasks]: Failed to fetch tasks, error: %v\n", err)
		return nil, err
	}
	return tasks, nil
}

func GetTaskById(id uint) (task models.Task, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&task).Error
	if err != nil {
		logger.Error.Printf("[repository.GetTaskById]: Failed to fetch task, ID: %v, error: %v\n", id, err)
		return task, err
	}
	return task, nil
}

func AddTask(task models.Task) error {
	result := db.GetDBConn().Create(&task)
	if result.Error != nil {
		logger.Error.Printf("[repository.AddTask]: Failed to add task, error: %v\n", result.Error)
		return result.Error
	}
	return nil
}

func UpdateTask(taskID uint, title, description string) error {
	var task models.Task
	result := db.GetDBConn().Model(&task).Where(
		"id = ?", taskID).Updates(models.Task{Title: title, Description: description})
	if result.Error != nil {
		logger.Error.Printf("[repository.UpdateTask]: Failed to update task with ID %v, error: %v\n", taskID, result.Error)
		return result.Error
	}
	return nil
}

func ToggleStatus(taskID uint) error {
	var task models.Task
	result := db.GetDBConn().Model(&task).Where(
		"id = ?", taskID).Update("IsDone", !task.IsDone)
	if result.Error != nil {
		logger.Error.Printf("[repository.ToggleStatus]: Failed to toggle status for task with ID %v, error: %v\n", taskID, result.Error)
		return result.Error
	}
	return nil
}

func DeleteTask(id uint) error {
	var task models.Task
	result := db.GetDBConn().Model(&task).Where("id = ?", id).Update("is_deleted", true)
	if result.Error != nil {
		logger.Error.Printf("[repository.DeleteTask]: Failed to delete task with ID %v, error: %v\n", id, result.Error)
		return result.Error
	}
	return nil
}

func InsertExistingData(tasks []models.Task) error {
	for _, task := range tasks {
		result := db.GetDBConn().Create(&task)
		if result.Error != nil {
			logger.Error.Printf("[repository.InsertExistingData]: Failed to insert task, error: %v\n", result.Error)
			return result.Error
		}
	}
	return nil
}

func SetPriority(taskID uint, priority int) error {
	var task models.Task
	result := db.GetDBConn().Model(&task).Where(
		"id = ?", taskID).Update("Priority", priority)
	if result.Error != nil {
		logger.Error.Printf("[repository.SetPriority]: Failed to set priority for task with ID %v, error: %v\n", taskID, result.Error)
		return result.Error
	}
	return nil
}

func GetTasksByIsDone(status string) ([]models.Task, error) {
	var tasks []models.Task
	result := db.GetDBConn().Where("is_done = ?", status == "completed").Find(&tasks)
	if result.Error != nil {
		logger.Error.Printf("[repository.GetTasksByIsDone]: Failed to fetch tasks, error: %v\n", result.Error)
		return nil, result.Error
	}
	return tasks, nil
}

func SortTasksByDate() ([]models.Task, error) {
	var tasks []models.Task
	result := db.GetDBConn().Order("created_at").Find(&tasks)
	if result.Error != nil {
		logger.Error.Printf("[repository.SortTasksByDate]: Failed to sort tasks, error: %v\n", result.Error)
		return nil, result.Error
	}
	return tasks, nil
}

func SortTasksByStatus() ([]models.Task, error) {
	var tasks []models.Task
	result := db.GetDBConn().Order("is_done").Find(&tasks)
	if result.Error != nil {
		logger.Error.Printf("[repository.SortTasksByStatus]: Failed to sort tasks, error: %v\n", result.Error)
		return nil, result.Error
	}
	return tasks, nil
}

func SortTasksByPriority() ([]models.Task, error) {
	var tasks []models.Task
	result := db.GetDBConn().Order("Priority").Find(&tasks)
	if result.Error != nil {
		logger.Error.Printf("[repository.SortTasksByPriority]: Failed to sort tasks, error: %v\n", result.Error)
		return nil, result.Error
	}
	return tasks, nil
}
