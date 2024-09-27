package main

import "fmt"

func main() {
	m := map[string]int{
		"james":      42,
		"Moneypenny": 43,
	}
	for key, value := range m {
		fmt.Printf("at key %v the value is %v\n", key, value)
	}
}
