package main

import (
	"fmt"
)

func main() {
	var todoList = [3]string{"купить хлеб"}

	todoList[1] = "купить молоко"
	todoList[2] = "купить пиво"

	fmt.Printf("К-во элементов в списке: %d\n", len(todoList))

	for _, item := range todoList {
		fmt.Printf("%s\n", item)
	}
}
