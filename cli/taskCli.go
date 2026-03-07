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
const ERROR_OPENING_FILE = "Error opening file: "
const ERROR_CREATING_FILE = "Error creating file: "
const ERROR_MARSHALLING = "Error marshalling JSON: "
const ERROR_WRITING_TO_FILE = "Error writing to file: "
const ERROR_FILE_INTERACTION = "Error reading from file: "

func AddTask(description Description) bool {
	// Чтение файла
	file, err := os.OpenFile(FILE_NAME, os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(ERROR_OPENING_FILE, err)
		file, err = os.Create(FILE_NAME)
		if err != nil {
			fmt.Println(ERROR_CREATING_FILE, err)
			return false
		}
		defer file.Close()
	}

	taskList, err := readTasks()

	// Добавление новой задачи
	task := Task{
		Id:          Id(len(taskList) + 1),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: description,
		TaskStatus:  Statuses.Todo,
	}
	taskList = append(taskList, task)
	tasksFile := TasksFile{taskList}

	// Парсинг в JSON
	jsonTask, err := json.MarshalIndent(tasksFile, "", "  ")
	if err != nil {
		fmt.Println(ERROR_MARSHALLING, err)
		return false
	}

	// Запись в файл
	_, err = file.Write(jsonTask)
	if err != nil {
		fmt.Println(ERROR_WRITING_TO_FILE, err)
		return false
	}

	fmt.Printf("Task added successfully (ID: %d)\n", task.Id)
	return true
}

// GetTasks получает все задачи независимо от статуса
func GetTasks() {
	tasks, err := readTasks()
	if err != nil {
		fmt.Println(ERROR_FILE_INTERACTION, err)
	}

	for _, task := range tasks {
		fmt.Printf(
			"Задача:\n - id=%d;\n - описание=\"%s\";\n - статус=%s;\n - время создания=%s;\n - последнее обновление=%s",
			task.Id,
			task.Description,
			task.TaskStatus,
			task.CreatedAt,
			task.UpdatedAt)
	}
}

// readTasks вспомогательная функция для чтения задач из файла
func readTasks() ([]Task, error) {
	data, err := os.ReadFile(FILE_NAME)
	if err != nil {
		return nil, err
	}

	var tasksFile TasksFile
	err = json.Unmarshal(data, &tasksFile)
	if err != nil {
		return nil, err
	}

	if tasksFile.Tasks == nil {
		tasksFile.Tasks = []Task{}
	}

	return tasksFile.Tasks, nil
}
