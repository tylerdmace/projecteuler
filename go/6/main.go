package main

import "fmt"

func main() {
	var sum int
	var sumOfSquares int
	var squareOfSums int

	for i := 1; i <= 100; i++ {
		sumOfSquares += i * i
	}

	for j := 1; j <= 100; j++ {
		sum += j
	}

	squareOfSums = sum * sum

	fmt.Printf("Difference: %v\r\n", squareOfSums-sumOfSquares)
}
