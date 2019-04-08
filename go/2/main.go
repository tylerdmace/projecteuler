package main

import "fmt"

func main() {
	max := 4000000
	a := 0
	b := 1

	fmt.Println(fibSum(a, b, 0, max))
}

func fibSum(a, b, sum, max int) int {
	if a+b > max {
		return sum
	} else {
		if (a+b)%2 == 0 {
			sum += a + b
		}

		return fibSum(b, a+b, sum, max)
	}
}
