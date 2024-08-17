package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/MadsSRasmussen/go-todo/internal/manager"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: todo [add|list|remove] [task]")
		return
	}

	// Manager setup
	manager := manager.New("tasks.csv")
	defer manager.WriteToFile()

	command := os.Args[1]
	switch command {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("You must specify a task to add.")
			return
		}

		description := os.Args[2]
		manager.AddTask(description)

	case "list":

		manager.PrintTasks()

	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("You must provide an ID")
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task id")
			return
		}
		fmt.Println("Marking task as done")
		manager.RemoveTask(id)
	default:
		fmt.Printf("Unknown command %s\n", os.Args[1])
	}

}
