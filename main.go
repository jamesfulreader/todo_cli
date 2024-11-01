package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
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
	fmt.Println("enter 1 to add a task")
	fmt.Println("enter 2 to list out the current tasks")
	fmt.Println("enter 3 to complete a task")
	fmt.Println("enter 4 to delete a task")
}

func writeTaskToFile(task Task) {
	file, _ := os.OpenFile("tasks.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{
		strconv.Itoa(task.ID),
		task.Description,
		strconv.FormatBool(task.Completed),
		task.CreatedAt.Format("2006-01-02"),
		task.DueAt.Format("2006-01-02"),
	})
}

func addTask() {
	fmt.Print("\n")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter task description: ")
	description, _ := reader.ReadString('\n')
	fmt.Print("Enter due date (YYYY-MM-DD): ")
	dueDateString, _ := reader.ReadString('\n')
	dueDate, _ := time.Parse("2006-01-02", dueDateString[:len(dueDateString)-1])

	simpleID := rand.IntN(100)

	newTask := Task{
		ID:          simpleID,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		DueAt:       dueDate,
	}

	writeTaskToFile(newTask)
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

func main() {
	showMenu()

	var choice int
	fmt.Print("\nEnter your choice: ")
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
