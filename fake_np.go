package main

import (
	"fmt"
	"math"
)

func main() {
	var l, r int
	if _, err := fmt.Scanf("%d %d\n", &l, &r); err != nil {
		return
	}

	r_sqrt := int(math.Sqrt(float64(r)))
	not_prime := make([]bool, r_sqrt+1)
	for i := 2; i <= r_sqrt; i++ {
		for j := i; j <= r_sqrt; j++ {
			if i*j <= r_sqrt {
				not_prime[i*j] = true
			} else {
				break
			}
		}
	}
	primes := []int{2}
	for i := 3; i <= r_sqrt; i++ {
		if !not_prime[i] {
			primes = append(primes, i)
		}
	}

	if l != r || l%2 == 0 {
		fmt.Println(2)
	} else {
		for i := 0; i < len(primes); i++ {
			if l%primes[i] == 0 {
				fmt.Println(primes[i])
				return
			}
		}
		fmt.Println(l)
	}
}
