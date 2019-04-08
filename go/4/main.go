package main

import (
	"fmt"
	"strconv"
)

func main() {
	var l int

	for i := 1; i <= 999; i++ {
		for j := 1; j <= 999; j++ {
			var p int
			var r string

			p = i * j

			for _, c := range strconv.Itoa(p) {
				r = string(c) + r
			}

			if strconv.Itoa(p) == r {
				if p > l {
					l = p
				}
			}
		}
	}

	fmt.Printf("Largest product palindrome: %v\r\n", l)
}
