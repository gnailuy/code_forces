package main

import (
	"fmt"
)

func main() {
	var n, m int
	if _, err := fmt.Scanf("%d %d\n", &n, &m); err != nil {
		return
	}

	one_circle := sum_from_one(n)
	if m >= one_circle {
		m = m % one_circle
	}

	for i := 1; i <= n; i++ {
		if m >= i {
			m -= i
		} else {
			break
		}
	}

	fmt.Println(m)
}

func sum_from_one(n int) int {
	return (n * (n + 1)) / 2
}
