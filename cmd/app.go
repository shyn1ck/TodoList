package cmd

import (
	"fmt"
	"todoList/controllers"
)

func Run() {
	fmt.Println("Добро пожаловать в приложение для управления задачами!")
	for {
		fmt.Println("\nВыберите нужную команду:")
		fmt.Println("1. Просмотреть все задачи")
		fmt.Println("2. Добавить новую задачу")
		fmt.Println("3. Редактировать задачу")
		fmt.Println("4. Изменить статус задачи")
		fmt.Println("5. Удалить задачу")
		fmt.Println("6. Установить приоритет задачи")
		fmt.Println("7. Фильтровать задачи по статусу")
		fmt.Println("8. Отсортировать задачи")
		fmt.Println("9. Вставить тестовые задачи")
		fmt.Println("0. Выход")

		var cmd string
		fmt.Scan(&cmd)
		switch cmd {
		case "0":
			fmt.Println("Выход из программы. До свидания!")
			return
		case "1":
			controllers.GetAllTasks()
		case "2":
			controllers.AddTask()
		case "3":
			controllers.EditTask()
		case "4":
			controllers.ToggleTaskStatus()
		case "5":
			controllers.DeleteTask()
		case "6":
			controllers.SetTaskPriority()
		case "7":
			controllers.FilterTasksByIsDone()
		case "8":
			controllers.SortTasks()
		case "9":
			controllers.InsertDataTasks()
		default:
			fmt.Println("Неверная команда. Пожалуйста, попробуйте снова.")
		}
	}
}
