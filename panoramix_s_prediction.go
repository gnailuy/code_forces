package main

import (
	"fmt"
)

func main() {
	var n, m int
	if _, err := fmt.Scanf("%d %d\n", &n, &m); err != nil {
		return
	}

	numbers := make([]bool, m+1) // False for prime number
	numbers[1] = true
	numbers[2] = false
	numbers[3] = false
	for i := 2; i <= m; i++ {
		for j := 2; j <= m/i; j++ {
			numbers[i*j] = true
		}
	}

	if numbers[m] { // m is not prime
		fmt.Println("NO")
	} else {
		for i := n + 1; i < m; i++ {
			if !numbers[i] {
				fmt.Println("NO")
				return
			}
		}
		fmt.Println("YES")
	}
}
