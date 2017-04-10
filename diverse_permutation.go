package main

import (
	"fmt"
)

func main() {
	var n, k int
	if _, err := fmt.Scanf("%d %d\n", &n, &k); err != nil {
		return
	}

	last := 1
	for i := k; i > 0; i-- {
		fmt.Printf("%d ", last)
		if last <= (k+1)/2 {
			last = last + i
		} else {
			last = last - i
		}
	}
	fmt.Printf("%d ", last)
	for i := k + 2; i <= n; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
