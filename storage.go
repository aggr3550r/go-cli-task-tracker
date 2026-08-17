package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Fetches tasks from a file with the given filename. Missing file means empty task list
func FetchTasks(filename string) ([]Task, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, fmt.Errorf("read tasks from file %q: %w", filename, err)
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("decode tasks from file %q: %w", filename, err)
	}

	if tasks == nil {
		tasks = []Task{}
	}

	return tasks, nil
}

// Saves tasks to a file
func SaveTasks(filename string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")

	if err != nil {
		return fmt.Errorf("encode tasks %w", err)
	}

	data = append(data, '\n')

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("write tasks to file %q: %w", filename, err)
	}

	return nil
}
