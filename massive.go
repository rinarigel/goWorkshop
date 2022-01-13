package main

import (
	"fmt"
)

func main() {
	var todoList = [...]string{
		"купить хлеб",
		"купить то-то",
		"сделать то-то",
		"сходить туда-то",
	}
	fmt.Printf("Кол-во элементов в списке: %d\n", len(todoList))

	for index, item := range todoList {
		fmt.Printf("%d. %s\n", index, item)
	}
}
