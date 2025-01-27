package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID          int
	Description string
}

var tasks []Task
var nextID int

func addTask(scanner *bufio.Scanner) {
	fmt.Println("Add Description")
	scanner.Scan()
	desc := scanner.Text()
	nextID++
	task := Task{ID: 1, Description: "something1"}
	tasks = append(tasks, task)
	task1 := Task{ID: 2, Description: "something1"}
	tasks = append(tasks, task1)
	task2 := Task{ID: 3, Description: "something1"}
	tasks = append(tasks, task2)
	task3 := Task{ID: 4, Description: desc}
	tasks = append(tasks, task3)
	fmt.Println("Task Added")
}

func viewTask() {
	if len(tasks) == 0 {
		fmt.Println("Tasks are empty")
		return
	}
	fmt.Println("\nCurrent Tasks:")
	for _, task := range tasks {
		fmt.Printf("ID: %d, %s\n", task.ID, task.Description)
	}
}

func deleteTask(scanner *bufio.Scanner) {
	if len(tasks) == 0 {
		fmt.Println("Tasks are empty")
		return
	}
	fmt.Println("\nCurrent Tasks:")
	for _, task := range tasks {
		fmt.Printf("ID: %d, %s\n", task.ID, task.Description)
	}

	scanner.Scan()
	idStr := scanner.Text()
	choice, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println("Invalid ID, please ender a number")
		return
	}
	if choice > len(tasks) {
		fmt.Println("Invalid ID, out of range")
		return
	}

	for i, task := range tasks {
		if task.ID == choice {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nTo-Do List Application")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Delete Task")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addTask(scanner)
		case "2":
			viewTask()
		case "3":
			deleteTask(scanner)
		case "4":
			break
		default:
			fmt.Println("This command ddoes not exist")
		}
	}
}
