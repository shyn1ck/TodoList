package service

import (
	"errors"
	"todoList/models"
	"todoList/pkg/repository"
)

func GetAllTasks() (tasks []models.Task, err error) {
	tasks, err = repository.GetAllTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTaskByID(id uint) (task models.Task, err error) {
	task, err = repository.GetTaskById(id)
	if err != nil {
		return task, err
	}
	return task, nil
}

func CreateTask(task models.Task) error {
	if task.Title == "" {
		return errors.New("title cannot be empty")
	}
	err := repository.AddTask(task)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTask(taskID uint, title, description string) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}
	err := repository.UpdateTask(taskID, title, description)
	if err != nil {
		return err
	}
	return nil
}

func ToggleTaskStatus(taskID uint) error {
	err := repository.ToggleStatus(taskID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTask(taskID uint) error {
	err := repository.DeleteTask(taskID)
	if err != nil {
		return err
	}
	return nil
}

func InsertExistingTasks(tasks []models.Task) error {
	for _, task := range tasks {
		if task.Title == "" {
			return errors.New("task title cannot be empty")
		}
	}
	err := repository.InsertExistingData(tasks)
	if err != nil {
		return err
	}
	return nil
}

func SetTaskPriority(taskID uint, priority int) error {
	if priority < 0 {
		return errors.New("priority cannot be negative")
	}
	err := repository.SetPriority(taskID, priority)
	if err != nil {
		return err
	}
	return nil
}

func GetTasksByStatus(status string) (tasks []models.Task, err error) {
	tasks, err = repository.GetTasksByIsDone(status)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func SortTasksByDate() (tasks []models.Task, err error) {
	tasks, err = repository.SortTasksByDate()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func SortTasksByStatus() (tasks []models.Task, err error) {
	tasks, err = repository.SortTasksByStatus()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func SortTasksByPriority() (tasks []models.Task, err error) {
	tasks, err = repository.SortTasksByPriority()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
