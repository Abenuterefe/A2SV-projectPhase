package data

import (
	"A2SV-projectPhase/task4/models"
	"errors"
)

var tasks = []models.Task{}

func GetTasks() []models.Task {
	return tasks
}

func GetTaskByID(id string) (*models.Task, error) {
	for _, t := range tasks {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, errors.New("task not found")
}

func CreateTask(task models.Task) {
	tasks = append(tasks, task)
}

func UpdateTask(id string, updated models.Task) error {
	for i, t := range tasks {
		if t.ID == id {
			tasks[i] = updated
			return nil
		}
	}
	return errors.New("task not found")
}

func DeleteTask(id string) error {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
