package main

import (
	"fmt"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	diagonals := 1
	last := 1
	for i := 2; i < n; i += 2 {
		for j := 1; j <= 4; j++ {
			diagonals += last + i*j
		}
		last = last + i*4
	}

	fmt.Println(diagonals)
}
