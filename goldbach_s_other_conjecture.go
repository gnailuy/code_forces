package main

import (
	"fmt"
)

func main() {
	var max int
	if _, err := fmt.Scanf("%d\n", &max); err != nil {
		return
	}

	not_prime := make([]bool, max)
	not_prime[1] = true
	not_prime[2] = true
	not_prime[3] = false
	for i := 2; i < max; i++ {
		for j := i; j < max; j++ {
			if i*j < max {
				not_prime[i*j] = true
			} else {
				break
			}
		}
	}

	is_sps := make([]bool, max)
	for i := 1; i < max; i++ {
		if !not_prime[i] {
			for j := 1; ; j++ {
				sps := i + 2*j*j
				if sps < max {
					is_sps[sps] = true
				} else {
					break
				}
			}
		}
	}

	for i := 4; i < max; i++ {
		if i%2 != 0 && not_prime[i] && !is_sps[i] {
			fmt.Println(i)
			break
		}
	}
}
