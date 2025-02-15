package main

import (
	"fmt"
	
)

var taskIDCounter int = 1

type Task struct{
	ID int `json:"id"`
	Description string `json:"description"`
}

var tasks []Task

func addTask(description string){
	newTask := Task{
		ID:  taskIDCounter,
		Description: description,
	}
	tasks = append(tasks, newTask)
	taskIDCounter++

}

func printTasks(){
	for _, task := range tasks{
		fmt.Printf("ID: %d, Desc: %s\n", task.ID, task.Description)
	}
}

func main(){
	for {
		var description string
		fmt.Print("Введите описание задачи (или 'exit' для выхода): ")
		fmt.Scanln(&description)

		if description == "exit" {
			break
		}

		addTask(description)
		printTasks()
	}

	
}