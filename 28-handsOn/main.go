package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			continue
		}
		// print i when it is odd number
		fmt.Println(i)

	}
}
