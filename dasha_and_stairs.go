package main

import (
	"fmt"
)

func main() {
	var a, b int
	if _, err := fmt.Scanf("%d %d\n", &a, &b); err != nil {
		return
	}

	if abs(a-b) < 2 && a+b > 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func abs(n int) int {
	if n >= 0 {
		return n
	} else {
		return -n
	}
}
