package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 42; i++ {
		x := rand.Intn(5)
		fmt.Printf("Iteration Number is\t%v and value of x is\t%v \n", i, x)
	}

}
