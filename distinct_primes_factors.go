package main

import (
	"fmt"
	"math"
)

func main() {
	var max int
	if _, err := fmt.Scanf("%d\n", &max); err != nil {
		return
	}

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

	primes := make([]int, 0, int(max/2))
	for i := 2; i < max; i++ {
		if !not_prime[i] {
			primes = append(primes, i)
		}
	}

	composite_cache := make(map[int]map[int]int)

	start := -1
	idx := -1

	for i := 4; i < max; i++ {
		i_factors := make(map[int]int)
		current := i
		for not_prime[current] {
			if factors, ok := composite_cache[current]; ok {
				for k, v := range factors {
					i_factors[k] += v
				}
				break
			} else {
				upper := int(math.Sqrt(float64(current))) + 1
				for j := 0; j < len(primes) && primes[j] <= upper; j++ {
					if current%primes[j] == 0 {
						i_factors[primes[j]] += 1
						current /= primes[j]
						break
					}
				}
			}
		}
		if current != 1 && !not_prime[current] {
			i_factors[current] += 1
		}
		if current != i {
			composite_cache[i] = i_factors
		}

		if len(i_factors) == 4 {
			fmt.Println(i, i_factors)

			if start < 0 {
				start = i
				idx = 0
			} else if idx < 2 {
				idx += 1
			} else {
				fmt.Println(start)
				break
			}
		} else {
			start = -1
			idx = -1
		}
	}
}
