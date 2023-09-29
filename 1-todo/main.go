// TODO App

package main

import (
	"fmt"
	"os"
	"bufio"
)

type Task struct {
	ID   int
	Name string
}

var tasks []Task

func addTask(name string) {
	newID := len(tasks) + 1
	newTask := Task{
		ID: newID, 
		Name: name,
	}
	tasks = append(tasks, newTask)
}

func removeTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			// remove the task from the slice
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
}

func listTasks() {
	for _, task := range tasks {
		fmt.Printf("%d: %s\n", task.ID, task.Name)
	}
}

func main() {
	for {
		fmt.Println("TODO List:")
		listTasks()

		fmt.Println("Enter an option (add/remove/quit): ")
		var option string
		fmt.Scanln(&option)

		switch option {
		case "add":
			fmt.Println("Enter task name: ")
			var taskName string
			scanner := bufio.NewScanner(os.Stdin) // initiate bufioscanner
			scanner.Scan() // use `for scanner.Scan()` to keep reading
			taskName = scanner.Text()
			addTask(taskName)
		case "remove":
			fmt.Println("Enter task ID to remove: ")
			var taskID int
			fmt.Scanln(&taskID)
			removeTask(taskID)
		case "quit":
			os.Exit(0)
		default:
			fmt.Println("Invalid Option")
		}
	}
}
