package data

import (
	"sync"
	"todo_app/models"
)

var (
	tasks []models.Task
	todoMu sync.Mutex
)

func GetTasks() []models.Task {
	return tasks
}

func AddTask(task models.Task) {
	tasks = append(tasks, task)
}

func UpdateTask(id int, task string, done bool) {
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Task = task
			tasks[i].Done = done
		}
	}
}

func DeleteTask(id int) {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return
		}
	}
}