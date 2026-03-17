package main

import (
	"fmt"
	"os"
	"strconv"
	"task_cli/cli"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Use --help for more instructions")
		return
	}

	switch args[0] {
	case "--help":
		printUsage()
	case "add":
		if len(args) < 2 {
			fmt.Println("Error: Task description is required")
			return
		}
		cli.AddTask(cli.Description(args[1]))
	case "update":
		if len(args) < 3 {
			fmt.Println("Error: ID and new description are required")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Error: ID must be a number. Got: '%s'\n", args[1])
			return
		}
		cli.UpdateTask(cli.Id(id), cli.Description(args[2]))
	case "delete":
		if len(args) < 2 {
			fmt.Println("Error: Task ID is required")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Error: ID must be a number. Got: '%s'\n", args[1])
			return
		}
		cli.DeleteTask(cli.Id(id))
	case "list":
		cli.GetTasks("") // Все задачи
	case "list-done":
		cli.GetTasks(cli.Statuses.Done) // Только выполненные
	case "list-not-done":
		cli.GetTasks("not-done") // Все, кроме выполненных
	case "list-in-progress":
		cli.GetTasks(cli.Statuses.InProgress) // Только в процессе
	case "mark-in-progress":
		if len(args) < 2 {
			fmt.Println("Error: Task ID is required")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Error: ID must be a number. Got: '%s'\n", args[1])
			return
		}
		cli.MarkInProgress(cli.Id(id))
	case "mark-done":
		if len(args) < 2 {
			fmt.Println("Error: Task ID is required")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Error: ID must be a number. Got: '%s'\n", args[1])
			return
		}
		cli.MarkDone(cli.Id(id))
	default:
		fmt.Printf("'%s': command not found. Use --help for more instructions", args[0])
	}
}

func printUsage() {
	fmt.Println("Task cli options:")
	fmt.Println("\tadd <description>        		- Add new task")
	fmt.Println("\tupdate <ID> <description>		- Update task's description")
	fmt.Println("\tdelete <ID>                 		- Delete task")
	fmt.Println("\tlist                        		- List all tasks")
	fmt.Println("\tlist-done                   		- List all tasks that are done")
	fmt.Println("\tlist-not-done               		- List all tasks that are not done")
	fmt.Println("\tlist-in-progress            		- List all tasks that are in progress")
	fmt.Println("\tmark-in-progress <ID>       		- Change task's status to 'in-progress'")
	fmt.Println("\tmark-done <ID>              		- Change task's status to 'done'")
}
