package repository

import (
	"log"
	"todoList/db"
	"todoList/models"
)

func GetAllTasks() (tasks []models.Task, err error) {
	log.Println("repository.GetAllTasks: Fetching all tasks from the database")
	err = db.GetDBConn().Where("is_deleted = ?", false).Find(&tasks).Error
	if err != nil {
		log.Printf("repository.GetAllTasks: Failed to fetch tasks, error: %v\n", err)
		return nil, err
	}
	log.Println("repository.GetAllTasks: Successfully fetched all tasks")
	return tasks, nil
}

func GetTaskById(id uint) (task models.Task, err error) {
	log.Printf("repository.GetTaskById: Fetching task by ID %v from the database\n", id)
	err = db.GetDBConn().Where("id = ?", id).First(&task).Error
	if err != nil {
		log.Printf("repository.GetTaskById: Failed to fetch task, ID: %v, error: %v\n", id, err)
		return task, err
	}
	log.Printf("repository.GetTaskById: Successfully fetched task by ID %v\n", id)
	return task, nil
}

func AddTask(task models.Task) error {
	log.Println("repository.AddTask: Adding new task to the database")
	result := db.GetDBConn().Create(&task)
	if result.Error != nil {
		log.Printf("repository.AddTask: Failed to add task, error: %v\n", result.Error)
		return result.Error
	}
	log.Println("repository.AddTask: Task added successfully")
	return nil
}

func UpdateTask(taskID uint, title, description string) error {
	log.Printf("repository.UpdateTask: Updating task with ID %v\n", taskID)
	var task models.Task
	result := db.GetDBConn().Model(&task).Where(
		"id = ?", taskID).Updates(models.Task{Title: title, Description: description})
	if result.Error != nil {
		log.Printf("repository.UpdateTask: Failed to update task with ID %v, error: %v\n", taskID, result.Error)
		return result.Error
	}
	log.Printf("repository.UpdateTask: Task with ID %v updated successfully\n", taskID)
	return nil
}

func ToggleStatus(taskID uint) error {
	log.Printf("repository.ToggleStatus: Toggling status for task with ID %v\n", taskID)
	var task models.Task
	result := db.GetDBConn().Model(&task).Where(
		"id = ?", taskID).Update("IsDone", !task.IsDone)
	if result.Error != nil {
		log.Printf("repository.ToggleStatus: Failed to toggle status for task with ID %v, error: %v\n", taskID, result.Error)
		return result.Error
	}
	log.Printf("repository.ToggleStatus: Status toggled for task with ID %v successfully\n", taskID)
	return nil
}

func DeleteTask(id uint) error {
	log.Printf("repository.DeleteTask: Soft deleting task with ID %v\n", id)
	var task models.Task
	result := db.GetDBConn().Model(&task).Where("id = ?", id).Update("is_deleted", true)
	if result.Error != nil {
		log.Printf("repository.DeleteTask: Failed to delete task with ID %v, error: %v\n", id, result.Error)
		return result.Error
	}
	log.Printf("repository.DeleteTask: Task with ID %v deleted successfully\n", id)
	return nil
}

func InsertExistingData(tasks []models.Task) error {
	log.Println("repository.InsertExistingData: Inserting existing tasks into the database")
	for _, task := range tasks {
		result := db.GetDBConn().Create(&task)
		if result.Error != nil {
			log.Printf("repository.InsertExistingData: Failed to insert task, error: %v\n", result.Error)
			return result.Error
		}
	}
	log.Println("repository.InsertExistingData: All tasks inserted successfully")
	return nil
}

func SetPriority(taskID uint, priority int) error {
	log.Printf("repository.SetPriority: Setting priority %v for task with ID %v\n", priority, taskID)
	var task models.Task
	result := db.GetDBConn().Model(&task).Where(
		"id = ?", taskID).Update("Priority", priority)
	if result.Error != nil {
		log.Printf("repository.SetPriority: Failed to set priority for task with ID %v, error: %v\n", taskID, result.Error)
		return result.Error
	}
	log.Printf("repository.SetPriority: Priority for task with ID %v set successfully\n", taskID)
	return nil
}

func GetTasksByIsDone(status string) ([]models.Task, error) {
	log.Printf("repository.GetTasksByIsDone: Fetching tasks with status %v\n", status)
	var tasks []models.Task
	result := db.GetDBConn().Where("is_done = ?", status == "completed").Find(&tasks)
	if result.Error != nil {
		log.Printf("repository.GetTasksByIsDone: Failed to fetch tasks, error: %v\n", result.Error)
		return nil, result.Error
	}
	log.Printf("repository.GetTasksByIsDone: Successfully fetched tasks with status %v\n", status)
	return tasks, nil
}

func SortTasksByDate() ([]models.Task, error) {
	log.Println("repository.SortTasksByDate: Sorting tasks by date")
	var tasks []models.Task
	result := db.GetDBConn().Order("created_at").Find(&tasks)
	if result.Error != nil {
		log.Printf("repository.SortTasksByDate: Failed to sort tasks, error: %v\n", result.Error)
		return nil, result.Error
	}
	log.Println("repository.SortTasksByDate: Successfully sorted tasks by date")
	return tasks, nil
}

func SortTasksByStatus() ([]models.Task, error) {
	log.Println("repository.SortTasksByStatus: Sorting tasks by status")
	var tasks []models.Task
	result := db.GetDBConn().Order("is_done").Find(&tasks)
	if result.Error != nil {
		log.Printf("repository.SortTasksByStatus: Failed to sort tasks, error: %v\n", result.Error)
		return nil, result.Error
	}
	log.Println("repository.SortTasksByStatus: Successfully sorted tasks by status")
	return tasks, nil
}

func SortTasksByPriority() ([]models.Task, error) {
	log.Println("repository.SortTasksByPriority: Sorting tasks by priority")
	var tasks []models.Task
	result := db.GetDBConn().Order("Priority").Find(&tasks)
	if result.Error != nil {
		log.Printf("repository.SortTasksByPriority: Failed to sort tasks, error: %v\n", result.Error)
		return nil, result.Error
	}
	log.Println("repository.SortTasksByPriority: Successfully sorted tasks by priority")
	return tasks, nil
}
