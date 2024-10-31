package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int
	Description string
	Completed   bool
	CreatedAt   time.Time
	DueAt       time.Time
}

func showMenu() {
	fmt.Println("Welcome the the todo list app CLI")
	fmt.Println("Please select a number based on what you would like to do")
	fmt.Println("enter 1 to add")
	fmt.Println("enter 2 to list out the current tasks")
	fmt.Println("enter 3 to complete a task")
	fmt.Println("enter 4 to delete a task")
}

func main() {
	showMenu()

	var choice int
	fmt.Println("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		addTask()
	case 2:
		listTasks()
	case 3:
		markComplete()
	case 4:
		deleteTask()
	default:
		fmt.Println("expected a number 1 - 4")
		os.Exit(1)
	}
}

func addTask() {
	description := flag.String("description", "Please enter a description", "task description")
	flag.Parse()
	fmt.Println("add task:", *description)
}

func listTasks() {
	fmt.Println("Print tasks")
}

func markComplete() {
	fmt.Println("Mark Complete")
}

func deleteTask() {
	fmt.Println("Delete Task")
}
