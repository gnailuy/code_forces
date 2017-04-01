package main

import (
	"fmt"
)

func main() {
	const upper int = 28123

	abundant := make(map[int]int)
	for i := 1; i <= upper; i++ {
		if is_abundant(i) {
			abundant[i] = 1
		}
	}

	abundant_sums := make(map[int]int)
	for i, _ := range abundant {
		for j, _ := range abundant {
			abundant_sums[i+j] = 1
		}
	}

	sum := int64(0)
	for i := 1; i <= upper; i++ {
		if abundant_sums[i] <= 0 {
			sum += int64(i)
		}
	}

	fmt.Println(sum)
}

func is_abundant(n int) bool {
	sum := 0
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	if sum > n {
		return true
	} else {
		return false
	}
}
