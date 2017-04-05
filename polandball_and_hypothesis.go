package main

import (
	"fmt"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	if n > 1 && n%2 != 0 {
		fmt.Println(1)
	} else {
		for i := 2; i <= n*1000+1; i++ {
			for j := i; j <= n*1000+1; j++ {
				non_prime := i * j
				if (non_prime-1)%n == 0 {
					fmt.Println((non_prime - 1) / n)
					return
				}
			}
		}
	}
}
