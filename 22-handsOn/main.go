package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("X and Y are %v and %v", x, y)
	if x < 4 && y < 4 {
		fmt.Println("X and y are less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("X and y are greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("X is between 3 and 7 ")
	} else if y != 5 {
		fmt.Println("y Is not equals to 5")
	}

}
