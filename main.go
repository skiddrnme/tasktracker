package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var taskIDCounter int = 1

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Progress    bool      `json:"progress"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var tasks []Task

// Меню для выбора
func menu() int {
	var i int
	fmt.Println("\nВыберите опцию:")
	fmt.Println("1. Добавить задачу")
	fmt.Println("2. Просмотреть задачи")
	fmt.Println("3. Удалить задачу")
	fmt.Println("4. Обновить задачу")
	fmt.Println("5. Сохранить в JSON")
	fmt.Println("6. Выйти")
	fmt.Print("Ваш выбор: ")
	fmt.Scanln(&i)
	return i
}

// Добавление задачи
func addTask(description string) {
	newTask := Task{
		ID:          taskIDCounter,
		Description: description,
		Progress:    false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, newTask)
	taskIDCounter++
}

// Удаление задачи
func deleteTask(id int) {
	newTasks := []Task{}
	for _, task := range tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		}
	}
	tasks = newTasks
}

// Обновление задачи
func updateTask(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Progress = true
			tasks[i].UpdatedAt = time.Now()
			fmt.Println("Задача обновлена!")
			return
		}
	}
	fmt.Println("Задача не найдена!")
}

// Вывод списка задач
func printTasks() {
	if len(tasks) == 0 {
		fmt.Println("Список задач пуст!")
		return
	}
	for _, task := range tasks {
		fmt.Printf("ID: %d, Описание: %s, Прогресс: %t, Создано: %s, Обновлено: %s\n",
			task.ID, task.Description, task.Progress,
			task.CreatedAt.Format("2006-01-02 15:04:05"),
			task.UpdatedAt.Format("2006-01-02 15:04:05"))
	}
}

// Сохранение задач в JSON
func saveToFile() {
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Ошибка кодирования JSON:", err)
		return
	}

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
	}
	fmt.Println("Задачи сохранены в tasks.json")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		switch menu() {
		case 1:
			fmt.Print("Введите описание задачи: ")
			scanner.Scan()
			description := scanner.Text()
			if description == "" {
				fmt.Println("Описание не может быть пустым!")
				continue
			}
			addTask(description)

		case 2:
			printTasks()

		case 3:
			var id int
			fmt.Print("Введите ID задачи для удаления: ")
			fmt.Scanln(&id)
			deleteTask(id)

		case 4:
			var id int
			fmt.Print("Введите ID задачи для обновления: ")
			fmt.Scanln(&id)
			updateTask(id)

		case 5:
			saveToFile()

		case 6:
			fmt.Println("Выход из программы.")
			return

		default:
			fmt.Println("Некорректный ввод. Попробуйте снова.")
		}
	}
}
