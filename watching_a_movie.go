package main

import (
	"fmt"
)

func main() {
	var n, x int
	if _, err := fmt.Scanf("%d %d\n", &n, &x); err != nil {
		return
	}

	var watch int = 0

	var pos int = 1
	var s, e int
	for i := 0; i < n; i++ {
		if _, err := fmt.Scanf("%d %d\n", &s, &e); err != nil {
			return
		}

		remain := (s - pos) % x
		watch += (remain + e - s + 1)

		pos = e + 1
	}

	fmt.Println(watch)
}
