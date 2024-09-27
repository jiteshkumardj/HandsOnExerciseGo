package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("fir Loop and iteration is :", i)
		for j := 0; j < 5; j++ {
			fmt.Println("This is second loop and iteratin is ;", j)
		}
	}
}
