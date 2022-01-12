package main

import "fmt"

func main() {
	x := 10
	p := &x // p - указатель на переменную x, хранящий ее адрес

	fmt.Println("Значение x:", x)
	fmt.Println("Адрес х:", p)
	fmt.Println("Значение *p:", *p)

	*p = 15

	fmt.Println("Значение х после изменения *p:", x)
}
