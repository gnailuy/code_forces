package main

import (
	"fmt"
)

func main() {
	var a, b int
	if _, err := fmt.Scanf("%d %d\n", &a, &b); err != nil {
		return
	}

	year := 0
	for a <= b {
		a *= 3
		b *= 2
		year++
	}

	fmt.Println(year)
}
