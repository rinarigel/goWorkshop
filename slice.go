package main

import "fmt"

func main() {
	todoList := [4]string{
		"купить хлеб",
		"купить молоко",
		"купить пиво",
		"дописать третий модуль",
	}

	tasks := todoList[1:4]
	for i := range tasks {
		fmt.Println(tasks[i])
	}

	fmt.Println("----- ПОСЛЕ changeTasks() -----")
	changeTasks(tasks)

	for i := range tasks {
		fmt.Println(tasks[i])
	}
}

func changeTasks(tasks []string) {
	tasks[0] = "пройти курс по Go"
	tasks[1] = "сказать Максиму спасибо :)"
}
