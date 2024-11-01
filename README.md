# Go CLI To-Do List

This is a simple command-line to-do list application written in Go. It allows you to manage your tasks efficiently from your terminal.

## Features

* Add tasks with descriptions and due dates.
* List all tasks with their status (pending or completed).
* Mark tasks as complete.
* Delete tasks.

## Usage

1.  Clone the repository:
    ```bash
    git clone [invalid URL removed]
    ```

2.  Navigate to the project directory:
    ```bash
    cd go-cli-todo
    ```

3.  Build the executable:
    ```bash
    go build
    ```

4.  Run the application:
    ```bash
    ./main
    ```

5.  Follow the on-screen menu to manage your tasks.

## Menu Options

* **1:** Add a new task.
* **2:** List all tasks.
* **3:** Mark a task as complete.
* **4:** Delete a task.
* **0:** Quit the application.

## Data Storage

The to-do list data is stored in a CSV file named `tasks.csv` in the project directory.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
