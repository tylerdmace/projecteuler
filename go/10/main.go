package main

import (
	"fmt"
	"math/big"
)

func main() {
	sum := 0
	target := 2000000

	for i := 2; i <= target; i++ {
		n := big.NewInt((int64)(i))
		if n.ProbablyPrime(1) {
			sum += i
		}
	}

	fmt.Printf("Sum of all primes under %v: %v\r\n", target, sum)
}
