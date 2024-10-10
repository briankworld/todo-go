package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	TaskName  string
	Completed bool
}

var tasks []Task

func addTask(task string) {
	newTask := Task{
		TaskName:  task,
		Completed: false,
	}

	tasks = append(tasks, newTask)
	fmt.Println("Task added")
}

func listTasks() {
	for i, task := range tasks {
		status := "not yet"
		if task.Completed {
			status = "done"
		}

		fmt.Printf("%d. %s [%s]\n", i+1, task.TaskName, status)
	}
}

func markCompleted(index int) {
	if index >= 1 && index <= len(tasks) {
		tasks[index-1].Completed = true
		fmt.Println("Task marked as completed")
	} else {
		fmt.Println("Invalid index")
	}
}

func editTask(index int, newTask string) {
	if index >= 1 && index <= len(tasks) {
		tasks[index-1].TaskName = newTask
		fmt.Println("Task edited")
	} else {
		fmt.Println("Invalid index")
	}
}

func deleteTask(index int) {
	if index >= 1 && index <= len(tasks) {
		tasks = append(tasks[:index-1], tasks[index:]...)
		fmt.Println("Task deleted")
	} else {
		fmt.Println("Invalid index")
	}
}

func main() {
	var indexInput int
	var taskInput, newTaskInput string

	fmt.Println("1. Add task")
	fmt.Println("2. List tasks")
	fmt.Println("3. Mark as completed")
	fmt.Println("4. Edit task")
	fmt.Println("5. Delete task")
	fmt.Println("6. Exit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter choice: ")
		scanner.Scan()
		input := scanner.Text()

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid choice")
		}

		switch choice {
		case 1:
			fmt.Print("Enter task: ")
			scanner.Scan()
			taskInput = scanner.Text()
			addTask(taskInput)
		case 2:
			listTasks()
		case 3:
			fmt.Print("Enter index: ")
			scanner.Scan()
			indexInput, _ = strconv.Atoi(scanner.Text())
			markCompleted(indexInput)
		case 4:
			fmt.Print("Enter index: ")
			scanner.Scan()
			indexInput, _ = strconv.Atoi(scanner.Text())
			fmt.Print("Enter task: ")
			scanner.Scan()
			newTaskInput = scanner.Text()
			editTask(indexInput, newTaskInput)
		case 5:
			fmt.Print("Enter index: ")
			scanner.Scan()
			indexInput, _ = strconv.Atoi(scanner.Text())
			deleteTask(indexInput)
		case 6:
			os.Exit(0)
		default:
			fmt.Println("Invaoid choice")
		}
	}
}
