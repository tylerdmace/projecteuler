package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	var discovered int
	nthPrime := 10001

	for i := 2; i > 0; i++ {
		n := big.NewInt((int64)(i))
		if n.ProbablyPrime(1) {
			discovered++

			if discovered == nthPrime {
				fmt.Printf("%vth prime: %v\r\n", nthPrime, i)
				os.Exit(0)
			}
		}
	}
}
