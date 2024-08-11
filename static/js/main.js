const taskInput = document.getElementById('taskInput');
const addTaskBtn = document.getElementById('addTaskBtn');
const taskList = document.getElementById('taskList');
const filterSelect = document.getElementById('filterSelect');
const sortSelect = document.getElementById('sortSelect');

async function fetchTasks(endpoint, method = 'GET', body = null) {
    try {
        const response = await fetch(endpoint, {
            method: method,
            headers: {
                'Content-Type': 'application/json'
            },
            body: body ? JSON.stringify(body) : null
        });
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return await response.json();
    } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
    }
}

// Fetch initial tasks
async function loadTasks() {
    const tasks = await fetchTasks('/tasks'); // Assuming your backend API endpoint is /tasks
    renderTasks(tasks);
}

// Render tasks to the list
function renderTasks(tasks) {
    taskList.innerHTML = '';
    tasks.forEach(task => {
        const li = document.createElement('li');
        li.setAttribute('data-task-id', task.id);

        const checkbox = document.createElement('input');
        checkbox.type = 'checkbox';
        checkbox.classList.add('task-checkbox');
        checkbox.checked = task.completed;
        checkbox.addEventListener('change', () => {
            toggleTaskStatus(task.id);
        });

        const taskContent = document.createElement('span');
        taskContent.classList.add('task-content');
        taskContent.textContent = task.text;
        if (task.completed) {
            taskContent.classList.add('completed');
        }

        const actions = document.createElement('div');
        actions.classList.add('task-actions');

        const editBtn = document.createElement('button');
        editBtn.classList.add('edit-btn');
        editBtn.textContent = 'Edit';
        editBtn.addEventListener('click', () => {
            editTask(task.id);
        });

        const deleteBtn = document.createElement('button');
        deleteBtn.classList.add('delete-btn');
        deleteBtn.textContent = 'Delete';
        deleteBtn.addEventListener('click', () => {
            deleteTask(task.id);
        });

        actions.appendChild(editBtn);
        actions.appendChild(deleteBtn);

        li.appendChild(checkbox);
        li.appendChild(taskContent);
        li.appendChild(actions);

        taskList.appendChild(li);
    });
}

// Add task function
addTaskBtn.addEventListener('click', async () => {
    const taskText = taskInput.value.trim();
    if (taskText !== "") {
        const newTask = {
            text: taskText,
            completed: false,
            priority: "normal",
            dueDate: null
        };

        const addedTask = await fetchTasks('/tasks', 'POST', newTask);
        loadTasks(); // Refresh the list
        taskInput.value = "";
    }
});

// Toggle task status
async function toggleTaskStatus(taskId) {
    try {
        await fetchTasks(`/tasks/${taskId}/toggle`, 'PUT');
        loadTasks();
    } catch (error) {
        console.error('Error toggling task status:', error);
    }
}

// Edit task
async function editTask(taskId) {
    // 1. Find the task to edit in the DOM.
    const taskItem = taskList.querySelector(`li[data-task-id="${taskId}"]`);
    const taskContent = taskItem.querySelector('.task-content');
    const taskText = taskContent.textContent; // Get current text

    // 2. Create an input element to edit the task text
    const input = document.createElement('input');
    input.type = 'text';
    input.value = taskText;
    input.classList.add('edit-input');
    taskContent.textContent = '';
    taskContent.appendChild(input);

    // 3. Replace the edit button with a save button
    const editBtn = taskItem.querySelector('.edit-btn');
    const saveBtn = document.createElement('button');
    saveBtn.classList.add('save-btn');
    saveBtn.textContent = 'Save';
    editBtn.replaceWith(saveBtn);

    // 4. Handle the save button click
    saveBtn.addEventListener('click', async () => {
        const newTaskText = input.value;
        try {
            await fetchTasks(`/tasks/${taskId}`, 'PUT', { text: newTaskText });
            loadTasks(); // Refresh the list
        } catch (error) {
            console.error('Error updating task:', error);
        }
    });
}

// Delete task
async function deleteTask(taskId) {
    try {
        await fetchTasks(`/tasks/${taskId}`, 'DELETE');
        loadTasks();
    } catch (error) {
        console.error('Error deleting task:', error);
    }
}

// Filtering tasks
function filterTasks() {
    const filterValue = filterSelect.value;
    if (filterValue === 'all') {
        return fetchTasks('/tasks');
    } else if (filterValue === 'completed') {
        return fetchTasks('/tasks/filter?isDone=true');
    } else if (filterValue === 'pending') {
        return fetchTasks('/tasks/filter?isDone=false');
    }
}

// Sorting tasks (you can implement different sorting logic here)
async function sortTasks(tasksToSort) {
    const sortValue = sortSelect.value;
    if (sortValue === 'priority') {
        // Implement priority-based sorting
        return tasksToSort.sort((a, b) => {
            // You'll need to add priority logic to your tasks
            // For example:
            if (a.priority === 'high' && b.priority !== 'high') {
                return -1; // a comes before b
            } else if (a.priority !== 'high' && b.priority === 'high') {
                return 1; // b comes before a
            } else {
                return 0; // default order
            }
        });
    } else if (sortValue === 'dueDate') {
        // Implement due date-based sorting
        return tasksToSort.sort((a, b) => {
            if (a.dueDate && b.dueDate) {
                return new Date(a.dueDate) - new Date(b.dueDate);
            } else if (a.dueDate) {
                return -1;
            } else if (b.dueDate) {
                return 1;
            } else {
                return 0;
            }
        });
    }
    return tasksToSort;
}

// Event listeners for filter and sort changes
filterSelect.addEventListener('change', async () => {
    const filteredTasks = await filterTasks();
    renderTasks(filteredTasks);
});

sortSelect.addEventListener('change', async () => {
    const filteredTasks = await filterTasks();
    const sortedTasks = await sortTasks(filteredTasks);
    renderTasks(sortedTasks);
});

// Initial rendering of tasks
loadTasks();