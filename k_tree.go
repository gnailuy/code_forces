package main

import (
	"fmt"
)

func main() {
	var n, k, d int
	if _, err := fmt.Scanf("%d %d %d\n", &n, &k, &d); err != nil {
		return
	}

	const modulo int = 1000000007

	cache := make(map[string]int)
	ways := find_ways(n, 0, k, d, modulo, cache)

	fmt.Println(ways)
}

func find_ways(n, c, k, d, m int, cache map[string]int) int {
	cache_key := string(c) + string(d)

	if cache[cache_key] > 0 {
		return cache[cache_key]
	}

	if n-c < d {
		return 0
	}

	sum := 0
	for i := 1; i <= k; i++ {
		if c+i == n {
			if d == 0 || i >= d {
				sum += 1
				if sum == m {
					sum = 0
				}
			}
		} else if c+i < n {
			if i >= d {
				sum += find_ways(n, c+i, k, 0, m, cache)
			} else {
				sum += find_ways(n, c+i, k, d, m, cache)
			}
			if sum >= m {
				sum %= m
			}
		}
	}

	cache[cache_key] = sum
	return sum
}
