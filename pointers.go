package main

import "fmt"

func main() {
	x := 10
	fmt.Println("Значение x:", x)

	increment(&x)
	fmt.Println("Значение х после вызова функции inctement()", x)
}

func increment(p *int) {
	*p += 1
}
