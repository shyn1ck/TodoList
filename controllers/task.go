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
		fmt.Println("Ошибка при получении задач:", err)
		return
	}

	fmt.Println("Задачи:")
	for _, task := range tasks {
		status := "Не выполнена"
		if task.IsDone {
			status = "Выполнена"
		}
		fmt.Printf("ID: %d, Заголовок: %s, Описание: %s, Статус: %s, Приоритет: %d, Создано: %s\n",
			task.ID, task.Title, task.Description, status, task.Priority, task.CreatedAt)
	}
}

func EditTask() {
	var taskID uint
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
	var taskID uint

	fmt.Println("Введите ID задачи для изменения статуса:")
	_, err := fmt.Scan(&taskID)
	if err != nil {
		fmt.Println("Ошибка при чтении ID задачи:", err)
		return
	}

	err = repository.ToggleStatus(taskID)
	if err != nil {
		fmt.Println("Ошибка при изменении статуса задачи:", err)
		return
	}

	fmt.Println("Статус задачи успешно изменен!")
}

func DeleteTask() {
	var taskID uint

	fmt.Println("Введите ID задачи для удаления:")
	_, err := fmt.Scan(&taskID)
	if err != nil {
		fmt.Println("Ошибка при чтении ID задачи:", err)
		return
	}

	err = repository.DeleteTask(taskID)
	if err != nil {
		fmt.Println("Ошибка при удалении задачи:", err)
		return
	}

	fmt.Println("Задача успешно удалена!")
}

func InsertDataTasks() {

	tasks := []models.Task{
		{Title: "Изучить основы Go", Description: "Пройти курс по основам Go на платформе онлайн-обучения", Priority: 1},
		{Title: "Сделать домашнее задание по алгоритмам", Description: "Решить задачи по алгоритмам и структурам данных", Priority: 2},
		{Title: "Работа с базами данных", Description: "Изучить SQL и ORM для работы с базами данных", Priority: 3},
		{Title: "Разработка API", Description: "Разработать REST API для учебного проекта", Priority: 2},
		{Title: "Прочитать книгу по программированию", Description: "Прочитать книгу 'Чистый код' Роберта Мартина", Priority: 1},
		{Title: "Участие в хакатоне", Description: "Принять участие в студенческом хакатоне", Priority: 3},
		{Title: "Создать GitHub репозиторий", Description: "Создать и настроить репозиторий для учебных проектов на GitHub", Priority: 2},
		{Title: "Установить среду разработки", Description: "Настроить IDE и все необходимые плагины для работы", Priority: 1},
		{Title: "Проектирование системы", Description: "Разработать архитектуру системы для учебного проекта", Priority: 3},
		{Title: "Учебный проект: веб-приложение", Description: "Создать простое веб-приложение с использованием фреймворка", Priority: 2},
	}

	err := repository.InsertExistingData(tasks)
	if err != nil {
		fmt.Println("Ошибка при вставке тестовых задач:", err)
		return
	}
	fmt.Println("Тестовые задачи успешно вставлены!")
}

func SetTaskPriority() {
	var taskID uint
	var priority int

	fmt.Println("Введите ID задачи для установки приоритета:")
	_, err := fmt.Scan(&taskID)
	if err != nil {
		fmt.Println("Ошибка при чтении ID задачи:", err)
		return
	}

	fmt.Println("Введите уровень приоритета (1-5):")
	_, err = fmt.Scan(&priority)
	if err != nil {
		fmt.Println("Ошибка при чтении уровня приоритета:", err)
		return
	}

	err = repository.SetPriority(taskID, priority)
	if err != nil {
		fmt.Println("Ошибка при установке приоритета задачи:", err)
		return
	}

	fmt.Println("Приоритет задачи успешно установлен!")
}

func FilterTasksByIsDone() {
	var status string

	fmt.Println("Введите статус для фильтрации (выполнена/не выполнена):")
	_, err := fmt.Scan(&status)
	if err != nil {
		fmt.Println("Ошибка при чтении статуса:", err)
		return
	}

	status = strings.ToLower(status)
	if status != "выполнена" && status != "не выполнена" {
		fmt.Println("Недопустимый статус. Пожалуйста, введите 'выполнена' или 'не выполнена'.")
		return
	}

	tasks, err := repository.GetTasksByIsDone(status)
	if err != nil {
		fmt.Println("Ошибка при фильтрации задач:", err)
		return
	}

	fmt.Println("Отфильтрованные задачи:")
	for _, task := range tasks {
		status := "Не выполнена"
		if task.IsDone {
			status = "Выполнена"
		}
		fmt.Printf("ID: %d, Заголовок: %s, Описание: %s, Статус: %s, Приоритет: %d, Создано: %s\n",
			task.ID, task.Title, task.Description, status, task.Priority, task.CreatedAt)
	}
}

func SortTasks() {
	var sortOption string

	fmt.Println("Введите опцию сортировки (data/status/priority):")
	_, err := fmt.Scan(&sortOption)
	if err != nil {
		fmt.Println("Ошибка при чтении опции сортировки:", err)
		return
	}

	var tasks []models.Task
	switch strings.ToLower(sortOption) {
	case "data":
		tasks, err = repository.SortTasksByDate()
	case "status":
		tasks, err = repository.SortTasksByStatus()
	case "priority":
		tasks, err = repository.SortTasksByPriority()
	default:
		fmt.Println("Недопустимая опция сортировки. Используйте 'data', 'status' или 'priority'.")
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
