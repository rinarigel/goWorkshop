package main

import (
	"fmt"
	"math"
)

func main() {
	exponentiation(10)
	reverse()
	reversTwoVariant()
}

func exponentiation(n int) {
	for i := 1; i <= n; i++ {
		result := math.Pow(2, float64(i))
		fmt.Println(result)
	}
}

func reverse() {
	myInts := []int{5, 2, 6, 3, 1, 4}
	fmt.Printf("Ints %v reversed: %v\n", myInts, reverseInts(myInts))
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}

func reversTwoVariant() {

	s := []int{5, 2, 6, 3, 1, 4}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	fmt.Println(s)
}
