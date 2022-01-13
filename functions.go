package main

import (
	"errors"
	"fmt"
)

func main() {
	printResult(0.6, 1, 1)
}

func printResult(length float32, width float32, height float32) {
	rectangleArea, err := calculatedRectangleArea(length, width)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Длина: %.2f см\n", length)
	fmt.Printf("Ширина: %.2f см\n", width)
	fmt.Println("Формула для подсчета периметра прямоугольника: S=2(length+width)")

	fmt.Printf("Периметр прямоугольника: %.2f cm. kv\n", rectangleArea)

	rectangleSquare, err := calculatedRectangleSquare(length, width)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Формула для подсчета площади прямоугольника: S=length*width")

	fmt.Printf("Площадь прямоугольника: %.2f cm. kv\n", rectangleSquare)

	triangleArea, err := calculatedTriangleArea(length, height)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Длина: %.2f см\n", length)
	fmt.Printf("Высота: %.2f см\n", height)
	fmt.Println("Формула для подсчета периметра треугольника: S=0.5*height*length")

	fmt.Printf("Периметр треугольника: %.2f cm. kv\n", triangleArea)

	triangleSquare, err := calculatedTriangleSquare(length, width)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Формула для подсчета площади треугольника: S=length*width*0.5")

	fmt.Printf("Площадь треугольника: %.2f cm. kv\n", triangleSquare)
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

func calculatedTriangleArea(length, height float32) (float32, error) {
	if (length <= 0) || (height <= 0) {
		return float32(0), errors.New("Sides can't be negative!")
	}

	return float32(length*height) * 0.5, nil
}

func calculatedTriangleSquare(length, width float32) (float32, error) {
	if (length <= 0) || (width <= 0) {
		return float32(0), errors.New("Sides can't be negative!")
	}

	return float32(length*width) * 0.5, nil
}
