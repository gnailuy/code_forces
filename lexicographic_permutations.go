package main

import (
	"fmt"
)

func main() {
	var n, s int
	if _, err := fmt.Scanf("%d %d\n", &n, &s); err != nil {
		return
	}

	n_factorial := make(map[int]int)
	n_factorial[0] = 1

	for i := 1; i <= s; i++ {
		n_factorial[i] = n_factorial[i-1] * i
	}

	if n > n_factorial[s] {
		fmt.Println(-1)
	}

	tokens := make([]int, s)
	for i := 0; i < s; i++ {
		tokens[i] = i
	}

	for i := s; i >= 1; i-- {
		m := n_factorial[i-1]
		for j := i; j >= 1; j-- {
			if n <= j*m && n > (j-1)*m {
				n -= (j - 1) * m
				p := 0
				for k := 0; k < s; k++ {
					if tokens[k] >= 0 {
						p++
						if p == j {
							fmt.Printf("%d", tokens[k])
							tokens[k] = -1
							break
						}
					}
				}
				break
			}
		}
	}
	fmt.Println()
}
