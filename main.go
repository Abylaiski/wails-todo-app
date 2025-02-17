package main

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)


var assets embed.FS


type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Priority  string `json:"priority"`
	DueDate   string `json:"due_date"`
}


type TodoApp struct {
	Tasks []*Task
	Mutex sync.Mutex
	File  string
}


func NewTodoApp(file string) *TodoApp {
	app := &TodoApp{
		Tasks: []*Task{},
		File:  file,
	}
	app.LoadTasks()
	return app
}


func (t *TodoApp) LoadTasks() {
	file, err := os.Open(t.File)
	if err != nil {
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&t.Tasks)
	if err != nil {
		fmt.Println("Ошибка при загрузке задач:", err)
	}
}


func (t *TodoApp) SaveTasks() error {
	file, err := os.Create(t.File)
	if err != nil {
		fmt.Printf("Ошибка при создании файла: %v\n", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(t.Tasks)
	if err != nil {
		fmt.Printf("Ошибка при кодировании задач: %v\n", err)
		return err
	}

	fmt.Printf("Задачи сохранены в файл: %s\n", t.File)
	return nil
}


func (t *TodoApp) AddTask(title, priority, dueDate string) error {
	if title == "" {
		return fmt.Errorf("название задачи не может быть пустым")
	}
	if priority != "Низкий" && priority != "Средний" && priority != "Высокий" {
		return fmt.Errorf("некорректный приоритет")
	}
	fmt.Printf("Добавление задачи: %s, %s, %s\n", title, priority, dueDate)

	t.Mutex.Lock()
	defer t.Mutex.Unlock()

	newTask := &Task{
		ID:        len(t.Tasks) + 1,
		Title:     title,
		Completed: false,
		Priority:  priority,
		DueDate:   dueDate,
	}

	t.Tasks = append(t.Tasks, newTask)
	fmt.Printf("Задача добавлена, всего задач: %d\n", len(t.Tasks))

	err := t.SaveTasks()
	if err != nil {
		fmt.Printf("Ошибка при сохранении: %v\n", err)
		return err
	}

	fmt.Println("Задача успешно сохранена")
	return nil
}


func (t *TodoApp) RemoveTask(id int) {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()

	newTasks := []*Task{}
	for _, task := range t.Tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		}
	}

	t.Tasks = newTasks
	t.SaveTasks()
}


func (t *TodoApp) ToggleTaskCompletion(id int) {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()

	for _, task := range t.Tasks {
		if task.ID == id {
			task.Completed = !task.Completed
			break
		}
	}

	t.SaveTasks()
}


func (t *TodoApp) GetTasks() []*Task {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	fmt.Printf("Список задач запрошен, всего: %d\n", len(t.Tasks))
	return t.Tasks
}


func (t *TodoApp) startup(ctx context.Context) {
	fmt.Println("Приложение запущено")
}

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Ошибка получения директории:", err)
		return
	}

	tasksFile := filepath.Join(currentDir, "tasks.json")
	fmt.Printf("Файл задач: %s\n", tasksFile)

	todoApp := NewTodoApp(tasksFile)
	err = wails.Run(&options.App{
		Title:            "Todo App",
		Width:            800,
		Height:           600,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup:        todoApp.startup,
		Bind: []interface{}{
			todoApp,
		},
	})
	if err != nil {
		fmt.Println("Ошибка запуска:", err)
	}
}
