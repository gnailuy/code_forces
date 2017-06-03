package main

import (
	"fmt"
)

func main() {
	var a, b int
	if _, err := fmt.Scanf("%d %d\n", &a, &b); err != nil {
		return
	}

	ab := [2]int{a, b}
	people := 0
	amount := 1
	for ab[people] >= amount {
		ab[people] -= amount
		people = (people + 1) % 2
		amount++
	}

	if people == 0 {
		fmt.Println("Vladik")
	} else {
		fmt.Println("Valera")
	}
}
