package main

import (
	"fmt"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	if n < 3 {
		fmt.Println(1)
		for i := 2; i <= n+1; i++ {
			fmt.Printf("1 ")
		}
		fmt.Println()
		return
	}

	sieve := make([]int, n+2)

	for i := int64(2); i <= int64(n+1); i++ {
		for j := i; j <= int64(n+1); j++ {
			if i*j <= int64(n+1) {
				sieve[int(i*j)] = 1
			} else {
				break
			}
		}
	}

	fmt.Println(2)
	for i := 2; i <= n+1; i++ {
		if sieve[i] == 0 { // Prime
			fmt.Printf("1 ")
		} else { // Composite
			fmt.Printf("2 ")
		}
	}
	fmt.Println()
}
