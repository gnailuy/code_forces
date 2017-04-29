package main

import (
	"fmt"
)

func main() {
	const max int = 1000

	is_not_prime := make([]bool, (2*max+1)*(2*max+1)+1)
	is_not_prime[0] = true
	is_not_prime[1] = true
	for i := 2; i <= 2*max+1; i++ {
		for j := i; j <= 2*max+1; j++ {
			is_not_prime[i*j] = true
		}
	}

	largest := 0
	la := 0
	lb := 0
	for a := 1 - max; a < max; a++ {
		for b := -max; b <= max; b++ {
			for n := 0; ; n++ {
				fn := n*n + a*n + b
				if fn > 0 {
					if is_not_prime[fn] {
						if n > largest {
							largest = n
							la = a
							lb = b
						}
						break
					}
				} else {
					break
				}
			}
		}
	}

	fmt.Println(largest)
	fmt.Println(la, lb, la*lb)
}
