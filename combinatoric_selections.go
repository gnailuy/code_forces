package main

import (
	"fmt"
)

func main() {
	var max int = 100
	var max_result int = 1000000
	var cnt int = 0

	for i := 2; i <= max; i++ {
		middle := i / 2
		for j := 1; j <= middle; j++ {
			c_i_j := c(i, j)
			if c_i_j > max_result {
				cnt += (i + 1 - j*2)
				break
			}
		}
	}

	fmt.Println(cnt)
}

func c(i, j int) int {
	result := 1

	devidors := make(map[int]bool)
	for k := i - j; k >= 1; k-- {
		devidors[k] = true
	}

	for k := i; k >= j+1; k-- {
		if d, ok := devidors[k]; ok {
			if d {
				devidors[k] = false
				continue
			}
		}
		result *= k
	}

	for k, v := range devidors {
		if v {
			result /= k
		}
	}

	return result
}
