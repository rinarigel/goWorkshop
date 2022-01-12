package main

import (
	"fmt"
)

func main() {
	printCircleArea()
}

func printCircleArea() {
	const pi = 3.1415
	circleRadius := 2 // circle radius
	// circle area
	circleArea := float32(circleRadius) * float32(circleRadius) * pi

	fmt.Printf("Circle radius: %d cm\n", circleRadius)
	fmt.Println("Formula for calculating area of circle: S=πr2")

	fmt.Printf("Circle area: %.2f cm. kv\n", circleArea)
}
