package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type TasksFile struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Id          Id          `json:"id"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Description Description `json:"description"`
	TaskStatus  TaskStatus  `json:"task_status"`
}

type Id int

// type CreatedAt time.Time
// type UpdatedAt time.Time
type Description string
type TaskStatus string
type TaskStatuses struct {
	Todo       TaskStatus
	InProgress TaskStatus
	Done       TaskStatus
}

var Statuses = TaskStatuses{
	Todo:       "todo",
	InProgress: "in-progress",
	Done:       "done",
}

const FILE_NAME = "tasks.json"
const ERROR_OPENING_OR_CREATING_FILE = "Error creating file: "
const ERROR_MARSHALLING = "Error marshalling JSON: "
const ERROR_WRITING_TO_FILE = "Error writing to file: "
const ERROR_READING_FILE = "Error reading from file: "

func AddTask(description Description) bool {
	f, err := os.OpenFile(FILE_NAME, os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(ERROR_OPENING_OR_CREATING_FILE, err)
		return false
	}
	defer f.Close()

	tasks := TasksFile{[]Task{
		{
			Id:          1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Description: description,
			TaskStatus:  Statuses.Todo,
		}}}

	jsonTask, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		fmt.Println(ERROR_MARSHALLING, err)
		return false
	}

	_, err = f.Write(jsonTask)
	if err != nil {
		fmt.Println(ERROR_WRITING_TO_FILE, err)
		return false
	}

	fmt.Printf("Task added successfully (ID: %d)\n", tasks.Tasks[0].Id)
	return true
}

func GetTasks() {
	data, err := os.ReadFile(FILE_NAME)
	if err != nil {
		fmt.Println(ERROR_OPENING_OR_CREATING_FILE)
	}

	var tasks TasksFile
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println(ERROR_READING_FILE, err)
	}

	for _, task := range tasks.Tasks {
		fmt.Printf(
			"Задача:\n - id=%d;\n - описание=\"%s\";\n - статус=%s;\n - время создания=%s;\n - последнее обновление=%s",
			task.Id,
			task.Description,
			task.TaskStatus,
			task.CreatedAt,
			task.UpdatedAt)
	}
}
