package main

import (
	"fmt"
)

func main() {
	const max int = 10000000
	sum := 0

	is_not_prime := make([]bool, max+1)
	is_not_prime[0] = true
	is_not_prime[1] = true
	for i := 2; i <= max; i++ {
		if i*i > max {
			break
		}
		for j := i; j <= max; j++ {
			if i*j > max {
				break
			}
			is_not_prime[i*j] = true
		}
	}

	for i := 10; i <= max; i++ {
		if !is_not_prime[i] {
			if is_truncatable(is_not_prime, i) {
				fmt.Println("Truncatable:", i)
				sum += i
			}
		}
	}

	fmt.Println()
	fmt.Println("SUM:", sum)
}

func is_truncatable(is_not_prime []bool, n int) bool {
	dig := 0
	c := n
	for c > 0 {
		if is_not_prime[c] {
			return false
		}
		c /= 10
		dig += 1
	}
	d := n
	for d > 0 {
		if is_not_prime[d] {
			return false
		}
		d %= power(10, dig-1)
		dig -= 1
	}
	return true
}

func power(base, p int) int {
	res := 1
	for p > 0 {
		res *= base
		p -= 1
	}
	return res
}
