package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Println("The value of x is :", x)
	if x <= 100 {
		fmt.Println("less than 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("the number is between 101 and 200")
	} else {
		fmt.Println("the number is between 201 and 250")
	}
}
