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

//go:embed frontend/dist
var assets embed.FS

<<<<<<< HEAD
// Task представляет задачу в списке дел.
=======
// Структура задачи
>>>>>>> bc98f00f3cd7a41de62f4487d9a81b77821876e8
type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Priority  string `json:"priority"`
	DueDate   string `json:"due_date"`
}

<<<<<<< HEAD
// TodoApp управляет списком задач и их сохранением.
=======
// Основная структура приложения
>>>>>>> bc98f00f3cd7a41de62f4487d9a81b77821876e8
type TodoApp struct {
	Tasks []*Task
	Mutex sync.Mutex
	File  string
}

<<<<<<< HEAD
// NewTodoApp создает экземпляр TodoApp и загружает задачи из файла.
=======
// Конструктор приложения
>>>>>>> bc98f00f3cd7a41de62f4487d9a81b77821876e8
func NewTodoApp(file string) *TodoApp {
	app := &TodoApp{
		Tasks: []*Task{},
		File:  file,
	}
	app.LoadTasks()
	return app
}

<<<<<<< HEAD
// LoadTasks загружает список задач из файла JSON.
=======
// Загрузка задач из файла
>>>>>>> bc98f00f3cd7a41de62f4487d9a81b77821876e8
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

<<<<<<< HEAD
// SaveTasks сохраняет текущий список задач в файл JSON.
=======
// Сохранение задач в файл
>>>>>>> bc98f00f3cd7a41de62f4487d9a81b77821876e8
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

<<<<<<< HEAD
	fmt.Printf("Задачи сохранены в файл: %s\n", t.File)
	return nil
}

// AddTask добавляет новую задачу в список.
=======
	fmt.Printf("Задачи успешно сохранены в файл: %s\n", t.File)
	return nil
}

// Добавление новой задачи
>>>>>>> bc98f00f3cd7a41de62f4487d9a81b77821876e8
func (t *TodoApp) AddTask(title, priority, dueDate string) error {
	if title == "" {
		return fmt.Errorf("название задачи не может быть пустым")
	}
	if priority != "Низкий" && priority != "Средний" && priority != "Высокий" {
		return fmt.Errorf("некорректный приоритет")
	}
<<<<<<< HEAD
	fmt.Printf("Добавление задачи: %s, %s, %s\n", title, priority, dueDate)
=======
	fmt.Printf("Получен запрос на добавление задачи: %s, %s, %s\n", title, priority, dueDate)
>>>>>>> bc98f00f3cd7a41de62f4487d9a81b77821876e8

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
<<<<<<< HEAD
	fmt.Printf("Задача добавлена, всего задач: %d\n", len(t.Tasks))

	err := t.SaveTasks()
	if err != nil {
		fmt.Printf("Ошибка при сохранении: %v\n", err)
=======
	fmt.Printf("Задача добавлена, текущее количество задач: %d\n", len(t.Tasks))

	err := t.SaveTasks()
	if err != nil {
		fmt.Printf("Ошибка при сохранении задач: %v\n", err)
>>>>>>> bc98f00f3cd7a41de62f4487d9a81b77821876e8
		return err
	}

	fmt.Println("Задача успешно сохранена")
	return nil
}

<<<<<<< HEAD
// RemoveTask удаляет задачу по ID.
=======
// Удаление задачи
>>>>>>> bc98f00f3cd7a41de62f4487d9a81b77821876e8
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

<<<<<<< HEAD
// ToggleTaskCompletion меняет статус выполнения задачи.
=======
// Переключение статуса задачи
>>>>>>> bc98f00f3cd7a41de62f4487d9a81b77821876e8
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

<<<<<<< HEAD
// GetTasks возвращает текущий список задач.
func (t *TodoApp) GetTasks() []*Task {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	fmt.Printf("Список задач запрошен, всего: %d\n", len(t.Tasks))
	return t.Tasks
}

// startup выполняется при запуске приложения.
=======
// Получение списка задач
func (t *TodoApp) GetTasks() []*Task {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	fmt.Printf("Запрошен список задач, количество: %d\n", len(t.Tasks))
	return t.Tasks
}

// Инициализация Wails-приложения
>>>>>>> bc98f00f3cd7a41de62f4487d9a81b77821876e8
func (t *TodoApp) startup(ctx context.Context) {
	fmt.Println("Приложение запущено")
}

func main() {
<<<<<<< HEAD
	// Получаем путь к файлу задач.
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Ошибка получения директории:", err)
=======
	// Используем абсолютный путь или путь относительно исполняемого файла
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
>>>>>>> bc98f00f3cd7a41de62f4487d9a81b77821876e8
		return
	}

	tasksFile := filepath.Join(currentDir, "tasks.json")
<<<<<<< HEAD
	fmt.Printf("Файл задач: %s\n", tasksFile)

	// Создаем приложение и запускаем Wails.
=======
	fmt.Printf("Путь к файлу задач: %s\n", tasksFile)

>>>>>>> bc98f00f3cd7a41de62f4487d9a81b77821876e8
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
<<<<<<< HEAD
		fmt.Println("Ошибка запуска:", err)
=======
		fmt.Println("Error:", err)
>>>>>>> bc98f00f3cd7a41de62f4487d9a81b77821876e8
	}
}
