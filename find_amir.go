package main

import (
	"fmt"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	c := 0
	for i := 1; i <= n/2; i++ {
		c += cost(i, n-i+1, n)
		if n-i != i {
			c += cost(n-i+1, i+1, n)
		}
	}

	fmt.Println(c)
}

func cost(i, j, n int) int {
	return (i + j) % (n + 1)
}
