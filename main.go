package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: task [add|list|done]")
		return
	}

	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("error loading tasks:", err)
	}
	cmd := os.Args[1]

	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("usage: task add \"task title\"")
			return
		}
		title := os.Args[2]
		id := len(tasks) + 1
		task := Task{
			ID:    id,
			Title: title,
			Done:  false,
		}
		tasks = append(tasks, task)
		saveTasks(tasks)
		fmt.Println("added task: ", title)
	case "list":
		if len(tasks) == 0 {
			fmt.Println("no tasks")
			return
		}

		for _, t := range tasks {
			status := " "
			if t.Done {
				status = "x"
			}
			fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Title)
		}
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("usage: task done <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Done = true
				saveTasks(tasks)
				fmt.Println("marked task done:", tasks[i].Title)
				return
			}
			fmt.Println("task not found")
		}
	default:
		fmt.Println("unknown command")
	}

}
