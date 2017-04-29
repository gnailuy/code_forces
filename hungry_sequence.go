package main

import (
	"fmt"
)

func main() {
	const max int = 10000000

	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	start := max - n + 1
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", start+i)
	}
	fmt.Printf("\n")
}
