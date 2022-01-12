package main

import (
	"fmt"
)

func main() {
	printCircleArea(2)
	printCircleArea(5)
}

func printCircleArea(radius int) {

	fmt.Printf("Circle radius: %d cm\n", radius)
	fmt.Println("Formula for calculating area of circle: S=πr2")

	circleArea := calculatedCircleArea(radius)
	fmt.Printf("Circle area: %.2f cm. kv\n", circleArea)
}

func calculatedCircleArea(radius int) float32 {
	const pi = 3.1415
	return float32(radius) * float32(radius) * pi
}
