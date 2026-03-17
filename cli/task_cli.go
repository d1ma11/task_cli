package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"task_cli/util"
	"time"
)

type TasksFile struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Id          Id             `json:"id"`
	CreatedAt   util.CreatedAt `json:"created_at"`
	UpdatedAt   util.UpdatedAt `json:"updated_at"`
	Description Description    `json:"description"`
	TaskStatus  TaskStatus     `json:"task_status"`
}

type Id int
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
const errorReadingFile = "Error reading file: "
const errorMarshalling = "Error marshalling JSON: "
const errorWritingToFile = "Error writing to file: "
const errorFileInteraction = "Error reading from file: "
const errorNoSuchTaskById = "There is no such task with id: "
const errorFileNotFount = "File is not found"

func AddTask(description Description) bool {
	tasksFile := &TasksFile{Tasks: make([]Task, 0)}
	tasksFile, err := readTasks()

	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println(errorReadingFile, err)
			return false
		}
	}

	// Добавление новой задачи
	task := Task{
		Id:          tasksFile.getMaxId() + 1,
		CreatedAt:   util.CreatedAt(time.Now()),
		UpdatedAt:   util.UpdatedAt(time.Now()),
		Description: description,
		TaskStatus:  Statuses.Todo,
	}

	tasksFile.Tasks = append(tasksFile.Tasks, task)

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

func UpdateTask(id Id, newDescription Description) bool {
	tasksFile, err := readTasks()

	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println(errorReadingFile, err)
			return false
		}
		fmt.Println(errorFileNotFount)
		return false
	}

	task := tasksFile.getTaskById(id)

	if task == nil {
		fmt.Println(errorNoSuchTaskById, id)
		return false
	}

	task.Description = newDescription
	task.UpdatedAt = util.UpdatedAt(time.Now())

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
	fmt.Printf("Task updated successfully (ID: %d)\n", task.Id)
	return true
}

func DeleteTask(id Id) bool {
	tasksFile, err := readTasks()

	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println(errorReadingFile, err)
			return false
		}
		fmt.Println(errorFileNotFount)
		return false
	}

	tasksFile.deleteTaskById(id)

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

	fmt.Printf("Task deleted successfully (ID: %d)\n", id)
	return true
}

func MarkInProgress(id Id) bool {
	return changeTaskStatus(id, Statuses.InProgress)
}

func MarkDone(id Id) bool {
	return changeTaskStatus(id, Statuses.Done)
}

// GetTasks получает все задачи независимо от статуса
func GetTasks(status TaskStatus) {
	tasks, err := readTasks()
	if err != nil {
		fmt.Println(errorFileInteraction, err)
		return
	}

	for _, task := range tasks.Tasks {
		// Определяем, нужно ли показывать задачу
		shouldShow := false
		if status == "" {
			// Пустой статус - показывать все задачи
			shouldShow = true
		} else if status == "not-done" {
			// Все задачи кроме выполненных
			shouldShow = task.TaskStatus != Statuses.Done
		} else {
			// Фильтр по конкретному статусу
			shouldShow = task.TaskStatus == status
		}

		if shouldShow {
			fmt.Printf(
				"\nЗадача:\n - id=%d;\n - описание=\"%s\";\n - статус=%s;\n - время создания=%s;\n - последнее обновление=%s\n",
				task.Id,
				task.Description,
				task.TaskStatus,
				time.Time(task.CreatedAt).Format(time.DateTime),
				time.Time(task.UpdatedAt).Format(time.DateTime),
			)
		}
	}
}

// readTasks вспомогательная функция для чтения задач из файла.
func readTasks() (*TasksFile, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var tasksFile *TasksFile
	err = json.Unmarshal(data, &tasksFile)
	if err != nil {
		return nil, err
	}

	if tasksFile.Tasks == nil {
		tasksFile.Tasks = []Task{}
	}

	return tasksFile, nil
}

// getMaxId вспомогательная функция для нахождения максимального идентификатора в списке задач
func (taskList *TasksFile) getMaxId() Id {
	maxId := Id(0)
	for _, task := range taskList.Tasks {
		if task.Id > maxId {
			maxId = task.Id
		}
	}
	return maxId
}

// getTaskById вспомогательная функция для нахождения задачи по его идентификатору
func (taskList *TasksFile) getTaskById(id Id) *Task {
	for i := 0; i < len(taskList.Tasks); i++ {
		if taskList.Tasks[i].Id == id {
			return &taskList.Tasks[i]
		}
	}
	return nil
}

// deleteTaskById вспомогательная функция для удаления задачи по его идентификатору
func (taskList *TasksFile) deleteTaskById(id Id) {
	newTaskList := make([]Task, 0)
	for i := 0; i < len(taskList.Tasks); i++ {
		if taskList.Tasks[i].Id != id {
			newTaskList = append(newTaskList, taskList.Tasks[i])
		}
	}
	taskList.Tasks = newTaskList
}

func changeTaskStatus(id Id, newStatus TaskStatus) bool {
	tasksFile, err := readTasks()

	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println(errorReadingFile, err)
			return false
		}
		fmt.Println(errorFileNotFount)
		return false
	}

	task := tasksFile.getTaskById(id)

	if task == nil {
		fmt.Println(errorNoSuchTaskById, id)
		return false
	}

	task.TaskStatus = newStatus
	task.UpdatedAt = util.UpdatedAt(time.Now())

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
	fmt.Printf("Task status changed to \"%s\" successfully (ID: %d)\n", newStatus, task.Id)
	return true
}
