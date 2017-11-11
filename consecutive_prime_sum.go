package main

import (
	"fmt"
)

func main() {
	var max int = 1000000 // A million

	not_prime := make([]bool, max)
	not_prime[1] = true
	not_prime[2] = false
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

	longest := 0
	length := 0

	for i := 1; i < max; i++ {
		if !not_prime[i] {
			sum := i
			num_p := 1
			for j := i + 1; j < max; j++ {
				if !not_prime[j] {
					sum += j
					num_p += 1
					if sum >= max {
						break
					} else if !not_prime[sum] {
						if num_p > length {
							longest = sum
							length = num_p
							fmt.Println(longest, length)
						}
					}
				}
			}
		}
	}

	fmt.Println(longest, length)
}
