package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

var taskFileName = "tasks.json"

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	// at this point the main function has given us only the "args" we actually need ->
	// add "Title"
	// list
	// complete TASK_ID
	// delete TASK_ID

	if len(args) == 0 {
		printUsage()
		return fmt.Errorf("missing command")
	}

	switch args[0] {
	case "add":
		if len(args) != 2 {
			return fmt.Errorf("Usage: go add \"Task title\"")
		}
		// fetch tasks from file
		tasks, err := FetchTasks(taskFileName)
		if err != nil {
			return err
		}

		// add the new task to the fetched tasks
		tasks, err = AddTask(tasks, args[1])
		if err != nil {
			return err
		}

		// save the updated task list to the file
		if err := SaveTasks(taskFileName, tasks); err != nil {
			return err
		}

		added := tasks[len(tasks)-1]
		fmt.Printf("Added task %d: %s\n", added.ID, added.Title)
		return nil

	case "list":
		if len(args) != 1 {
			return fmt.Errorf("Usage: go run . list")
		}

		// fetch tasks for storage files
		tasks, err := FetchTasks(taskFileName)
		if err != nil {
			return err
		}

		FormatTaskList(tasks)
		return nil

	case "complete":
		taskID, err := commandID(args, "complete")
		if err != nil {
			return err
		}

		if len(args) != 2 {
			return fmt.Errorf("Usage: go run . complete TASK_ID")
		}

		// fetch tasks from storage
		// recall if file doesn't exist we just return an empty list of Task
		tasks, err := FetchTasks(taskFileName)
		if err != nil {
			return err
		}

		// at this point we can make the task with given ID completed
		tasks, err = CompleteTask(tasks, taskID)
		if err != nil {
			return err
		}

		// save completed task
		if err = SaveTasks(taskFileName, tasks); err != nil {
			return err
		}
		fmt.Printf("Completed task %d\n", taskID)
		return nil

	case "delete":
		taskID, err := commandID(args, "delete")
		if err != nil {
			return err
		}
		if len(args) != 2 {
			return fmt.Errorf("Usage: go run . delete TASK_ID")
		}
		tasks, err := FetchTasks(taskFileName)
		if err != nil {
			return err
		}
		tasks, err = DeleteTask(tasks, taskID)
		if err != nil {
			return err
		}
		if err = SaveTasks(taskFileName, tasks); err != nil {
			return err
		}
		fmt.Printf("Deleted task %d\n", taskID)
		return nil

	default:
		return fmt.Errorf("Unknown command: %q", args[0])
	}
}

func commandID(args []string, command string) (int, error) {
	if len(args) != 2 {
		return 0, fmt.Errorf("Usage: go run . %s TASK_ID", command)
	}

	id, err := strconv.Atoi(args[1])
	if err != nil || id <= 0 {
		return 0, fmt.Errorf("task ID must be a positive integer")
	}

	return id, nil
}

// FormatTaskList prints all tasks in an aligned tabular format.
func FormatTaskList(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tSTATUS\tTITLE")
	for _, task := range tasks {
		status := "[ ]"
		if task.Completed {
			status = "[x]"
		}
		fmt.Fprintf(w, "%d\t%s\t%s\n", task.ID, status, task.Title)
	}
	w.Flush()
}

func printUsage() {
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, " go run . add \"Task title\"")
	fmt.Fprintln(os.Stderr, " go run . list")
	fmt.Fprintln(os.Stderr, " go run . complete TASK_ID")
	fmt.Fprintln(os.Stderr, " go run . delete TASK_ID")
}
