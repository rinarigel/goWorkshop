package main

import (
	"errors"
	"fmt"
)

func main() {
	printResult(0.6, 1)
}

func printResult(length float32, width float32) {
	rectangleArea, err := calculatedRectangleArea(length, width)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("length: %d cm\n", length)
	fmt.Printf("width: %d cm\n", width)
	fmt.Println("Formula for calculating area of rectangle: S=2(length+width)")

	fmt.Printf("Rectangle area: %.2f cm. kv\n", rectangleArea)

	rectangleSquare, err := calculatedRectangleSquare(length, width)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Formula for calculating square of rectangle: S=length*width")

	fmt.Printf("Rectangle square: %.2f cm. kv\n", rectangleSquare)
}

func calculatedRectangleArea(length, width float32) (float32, error) {
	if (length <= 0) || (width <= 0) {
		return float32(0), errors.New("Sides can't be negative!")
	}

	return float32(length+width) * 2, nil
}

func calculatedRectangleSquare(length, width float32) (float32, error) {
	if (length <= 0) || (width <= 0) {
		return float32(0), errors.New("Sides can't be negative!")
	}

	return float32(length * width), nil
}
