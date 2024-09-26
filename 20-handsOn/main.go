package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	switch {
	case x <= 100:
		fmt.Println("less than 100")
	case x >= 101 && x <= 200:
		fmt.Println("the number is between 101 and 200")
	default:
		fmt.Println("the number is between 201 and 250")
	}

}
