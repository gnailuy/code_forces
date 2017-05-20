package main

import (
	"fmt"
)

func main() {
	var p int
	if _, err := fmt.Scanf("%d\n", &p); err != nil {
		return
	}

	res := 0

	for l := 1; ; l++ {
		start := power(10, l)
		end := power(10, l+1) - 1
		max_j := power(9, p) * l

		if max_j < start {
			break
		}

		for i := start; i <= end; i++ {
			if i > max_j {
				break
			}
			j := to_power_p(i, p)
			if i == j {
				fmt.Println(i)
				res += i
			}
		}
	}

	fmt.Println(res)
}

func to_power_p(i, p int) int {
	res := 0
	for i > 0 {
		d := i % 10
		res += power(d, p)
		i = i / 10
	}
	return res
}

func power(n, p int) int {
	res := n
	for i := 1; i < p; i++ {
		res *= n
	}
	return res
}
