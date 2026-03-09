package main

import (
	"os"
	"strconv"
	"task_cli/cli"
)

func main() {
	//for i, val := range os.Args {
	//	fmt.Printf("Arguments #%d: %s\n", i, val)
	//}

	args := os.Args[1:]
	//fmt.Println(args)

	switch args[0] {
	case "add":
		cli.AddTask(cli.Description(args[1]))
	case "update":
		id, _ := strconv.Atoi(args[1])
		cli.UpdateTask(cli.Id(id), cli.Description(args[2]))
	case "delete":
		id, _ := strconv.Atoi(args[1])
		cli.DeleteTask(cli.Id(id))
	case "list":
		cli.GetTasks()
	case "mark-in-progress":
		id, _ := strconv.Atoi(args[1])
		cli.MarkInProgress(cli.Id(id))
	case "mark-done":
		id, _ := strconv.Atoi(args[1])
		cli.MarkDone(cli.Id(id))
	}
}
