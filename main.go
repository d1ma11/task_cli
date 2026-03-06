package main

import (
	"os"
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
	case "delete":
	case "list":
	case "mark-in-progress":
	case "mark-done":
	}
}
