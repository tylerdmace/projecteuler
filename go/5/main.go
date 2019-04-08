package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 2520; i > 0; i++ {
		for j := 1; j <= 20; j++ {
			if i%j != 0 {
				break
			}

			if j == 20 {
				fmt.Printf("Smallest positive number that can be evenly divided by numbers 1 to 20: %v\r\n", i)
				os.Exit(0)
			}
		}
	}
}
