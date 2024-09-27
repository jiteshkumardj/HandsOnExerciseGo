package main

import "fmt"

func main() {
	m := map[string]int{
		"james":      42,
		"Moneypenny": 43,
	}
	age := m["james"]
	fmt.Println(age)
	if value, ok := m["Q"]; !ok {
		fmt.Println("Age of Q:", value)
	}
}
