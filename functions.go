package main

import (
	"errors"
	"fmt"
)

func main() {
	printCircleArea(-2)
}

func printCircleArea(radius int) {
	circleArea, err := calculatedCircleArea(radius)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Circle radius: %d cm\n", radius)
	fmt.Println("Formula for calculating area of circle: S=πr2")

	fmt.Printf("Circle area: %.2f cm. kv\n", circleArea)
}

func calculatedCircleArea(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("Radius of circle can't be negative")
	}

	const pi = 3.1415
	return float32(radius) * float32(radius) * pi, nil
}
