package db

const (
	AddNewTaskQuery = `INSERT INTO tasks (title, description, priority, created_at) 
					   VALUES ($1, $2, $3, NOW());`

	GetAllTasksQuery = `SELECT id, title, description, is_done, priority, created_at
						FROM tasks
						WHERE is_deleted = FALSE;`

	UpdateTaskQuery = `UPDATE tasks
					   SET title = $1, description = $2
					   WHERE id = $3;`

	ToggleStatusQuery = `UPDATE tasks
						 SET is_done = NOT is_done
						 WHERE id = $1;`

	DeleteTaskQuery = `UPDATE tasks
					   SET is_deleted = TRUE
					   WHERE id = $1;`

	SetPriorityQuery = `UPDATE tasks
						SET priority = $1
						WHERE id = $2;`

	GetTasksByIsDoneQuery = `SELECT id, title, description, is_done, priority, created_at
					FROM tasks
					WHERE is_done = $1 AND is_deleted = FALSE;`

	SortByDateQuery = `SELECT id, title, description, is_done, priority, created_at
					   FROM tasks
					   WHERE is_deleted = FALSE
					   ORDER BY created_at;`

	SortByStatusQuery = `SELECT id, title, description, is_done, priority, created_at
						 FROM tasks
						 WHERE is_deleted = FALSE
						 ORDER BY is_done;`

	SortByPriorityQuery = `SELECT id, title, description, is_done, priority, created_at
						   FROM tasks
						   WHERE is_deleted = FALSE
						   ORDER BY priority;`
	InsertExistingData = `INSERT INTO tasks (title, description, is_done, is_deleted, priority, created_at) VALUES
					('Закончить учебник по Go', 'Пройти учебник по программированию на Go на сайте Go.dev.', FALSE, FALSE, 1, NOW()),
					('Сдать задание', 'Отправить задание по алгоритмам до дедлайна.', FALSE, FALSE, 2, NOW()),
					('Прочитать книгу', 'Прочитать книгу "Практическое программирование" для улучшения навыков.', FALSE, FALSE, 3, NOW()),
					('Исправить ошибки', 'Исправить ошибки в проекте по управлению задачами.', FALSE, FALSE, 1, NOW()),
					('Посетить мастер-класс', 'Принять участие в онлайн мастер-классе по продвинутым техникам Go.', FALSE, FALSE, 2, NOW()),
					('Обновить резюме', 'Обновить резюме с последними проектами и навыками.', FALSE, FALSE, 3, NOW()),
					('Подготовиться к собеседованию', 'Просмотреть часто задаваемые вопросы и попрактиковаться в решении задач.', FALSE, FALSE, 4, NOW()),
					('Изучить Git', 'Закончить учебник по Git для улучшения навыков контроля версий.', FALSE, FALSE, 1, NOW()),
					('Создать портфолио', 'Создать веб-сайт портфолио для демонстрации проектов и достижений.', FALSE, FALSE, 5, NOW()),
					('Участвовать в открытых проектах', 'Найти проект с открытым исходным кодом и внести свой вклад.', FALSE, FALSE, 3, NOW());
`
)
