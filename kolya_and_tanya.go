package main

import (
	"fmt"
)

const modulo int64 = 1000000000 + 7

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	const all_ways int64 = 3 * 3 * 3
	const sei_ways int64 = 7

	a := power_modulo(all_ways, n)
	b := power_modulo(sei_ways, n)
	if a < b {
		a += modulo
	}

	fmt.Println(a - b)
}

// Calculate w to the power of n, in modulo of 1000000007
func power_modulo(w int64, n int) int64 {

	rs := make([]int64, n+1)
	rs[1] = w % modulo

	for i := 2; i <= n; i++ {
		p := rs[i-1] * w
		if p >= modulo {
			p %= modulo
		}

		rs[i] = p
	}

	return rs[n]
}
