package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
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
	fmt.Println("enter 0 to exit")
}

func generateNextID() int {
	tasks, _ := readTasksFromFile()
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].ID + 1
}

func readTasksFromFile() ([]Task, error) {
	file, _ := os.Open("tasks.csv")
	defer file.Close()
	reader := csv.NewReader(file)
	lines, _ := reader.ReadAll()
	var tasks []Task
	for _, line := range lines {
		id, _ := strconv.Atoi(line[0])
		completed, _ := strconv.ParseBool(line[2])
		createdAt, _ := time.Parse("2006-01-02", line[3])
		dueAt, _ := time.Parse("2006-01-02", line[4])
		tasks = append(tasks, Task{
			ID:          id,
			Description: line[1],
			Completed:   completed,
			CreatedAt:   createdAt,
			DueAt:       dueAt,
		})
	}
	return tasks, nil
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

	newTask := Task{
		ID:          generateNextID(),
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		DueAt:       dueDate,
	}

	writeTaskToFile(newTask)
}

func listTasks() {
	tasks, err := readTasksFromFile()
	if err != nil {
		fmt.Println("error reading tasks: ", err)
	}

	if len(tasks) == 0 {
		fmt.Println("no tasks found")
	}

	for _, task := range tasks {
		status := "pending"
		if task.Completed {
			status = "completed"
		}
		fmt.Printf("ID: %d | Description: %s | Due Date: %s | Status: %s\n", task.ID, task.Description[:len(task.Description)-1], task.DueAt.Format("2006-01-02"), status)
	}
}

func markComplete() {
	listTasks()
	var choice int
	fmt.Print("\nWhich task would you like to complete? Please enter its ID number: ")
	fmt.Scan(&choice)

	tasks, err := readTasksFromFile()
	if err != nil {
		fmt.Println("error reading tasks: ", err)
	}

	found := false
	for i, task := range tasks {
		if task.ID == choice {
			tasks[i].Completed = true
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Task not found")
		return
	}

	file, err := os.Create("tasks.csv")
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, task := range tasks {
		err := writer.Write([]string{
			strconv.Itoa(task.ID),
			task.Description,
			strconv.FormatBool(task.Completed),
			task.CreatedAt.Format("2006-01-02"),
			task.DueAt.Format("2006-01-02"),
		})
		if err != nil {
			fmt.Println("error writing to CSV:", err)
			return
		}
	}
}

func deleteTask() {
	listTasks()
	var choice int
	fmt.Print("\nWhich task would you like to delete? Please enter its ID number: ")
	fmt.Scan(&choice)

	tasks, err := readTasksFromFile()
	if err != nil {
		fmt.Println("error reading tasks: ", err)
	}

	found := false
	for i, task := range tasks {
		if task.ID == choice {
			tasks = append(tasks[:i], tasks[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Task not found")
		return
	}

	file, err := os.Create("tasks.csv")
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, task := range tasks {
		err := writer.Write([]string{
			strconv.Itoa(task.ID),
			task.Description,
			strconv.FormatBool(task.Completed),
			task.CreatedAt.Format("2006-01-02"),
			task.DueAt.Format("2006-01-02"),
		})
		if err != nil {
			fmt.Println("error writing to CSV:", err)
			return
		}
	}
}

func main() {
	for {
		showMenu()

		var choice int
		fmt.Print("\nEnter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 0:
			fmt.Println("Exiting . . .")
			return
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
}
