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

const fileName = "tasks.json"
const errorOpeningFile = "Error opening file: "
const errorMarshalling = "Error marshalling JSON: "
const errorWritingToFile = "Error writing to file: "
const errorFileInteraction = "Error reading from file: "

func AddTask(description Description) bool {
	taskList, err := readTasks()

	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println(errorOpeningFile, err)
			return false
		}
	}

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
		fmt.Println(errorMarshalling, err)
		return false
	}

	// Запись в файл
	err = os.WriteFile(fileName, jsonTask, 0644)
	if err != nil {
		fmt.Println(errorWritingToFile, err)
		return false
	}

	fmt.Printf("Task added successfully (ID: %d)\n", task.Id)
	return true
}

// GetTasks получает все задачи независимо от статуса
func GetTasks() {
	tasks, err := readTasks()
	if err != nil {
		fmt.Println(errorFileInteraction, err)
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
	data, err := os.ReadFile(fileName)
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
