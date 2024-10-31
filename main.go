package main

import (
	"bufio"
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
	fmt.Println("enter 4 to delete a task\n")
}

func main() {
	showMenu()

	var choice int
	fmt.Print("Enter your choice: ")
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
	fmt.Print("\n")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter task description: ")
	description, _ := reader.ReadString('\n')
	fmt.Print("Enter due date (YYYY-MM-DD): ")
	dueDateString, _ := reader.ReadString('\n')
	dueDate, _ := time.Parse("2006-01-02", dueDateString[:len(dueDateString)-1])

	fmt.Printf(description)
	fmt.Printf(dueDate.String())
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
