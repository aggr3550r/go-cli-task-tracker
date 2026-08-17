package main

import (
	"fmt"
	"strings"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func NextTaskID(tasks []Task) int {
	highestTaskID := 0

	for _, task := range tasks {
		if task.ID > highestTaskID {
			highestTaskID = task.ID
		}
	}

	return highestTaskID + 1
}

func AddTask(tasks []Task, title string) ([]Task, error) {
	title = strings.TrimSpace(title)

	if title == "" {
		return nil, fmt.Errorf("task title can not be empty")
	}

	task := Task{
		ID:    NextTaskID(tasks),
		Title: title,
	}

	return append(tasks, task), nil
}

// CompleteTask marks the task with the given ID (if found) as completed
func CompleteTask(tasks []Task, id int) ([]Task, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID must be a positive integer")
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			return tasks, nil
		}
	}

	return nil, fmt.Errorf("Task with ID %d not found", id)
}

// DeleteTask removes the task with the given ID (if found)
func DeleteTask(tasks []Task, id int) ([]Task, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID must be a positive integer")
	}

	for i := range tasks {
		if tasks[i].ID == id {
			return append(tasks[:i], tasks[i+1:]...), nil
		}
	}

	return nil, fmt.Errorf("Task with ID %d not found", id)
}
