package main

import "fmt"

func main() {
	val := 0
	max := 1000

	for i := 1; i < max; i++ {
		if i%3 == 0 || i%5 == 0 {
			val = val + i
		}
	}

	fmt.Println(val)
}
