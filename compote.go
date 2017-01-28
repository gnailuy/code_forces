package main

import (
	"fmt"
)

func main() {
	var l int
	if _, err := fmt.Scanf("%d\n", &l); err != nil {
		return
	}
	var a int
	if _, err := fmt.Scanf("%d\n", &a); err != nil {
		return
	}
	var p int
	if _, err := fmt.Scanf("%d\n", &p); err != nil {
		return
	}

	a /= 2
	p /= 4

	fmt.Println(7 * min(l, min(a, p)))
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
