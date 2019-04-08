package main

import "fmt"

func main() {
	target := 1000

	for a := 3; a < target; a++ {
		for b := a + 1; b < target; b++ {
			c := 1000 - a - b

			if a*a+b*b == c*c {
				fmt.Printf("Product of %v * %v * %v = %v\r\n", a, b, c, a*b*c)
			}
		}
	}
}
