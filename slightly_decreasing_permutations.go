package main

import (
	"fmt"
)

func main() {
	var n, k int
	if _, err := fmt.Scanf("%d %d\n", &n, &k); err != nil {
		return
	}

	p := make([]int, n*2)
	for i := 0; i < n; i++ {
		p[i] = i + 1
	}
	s1 := 0       // Inclusive
	s2 := 2*n - 1 // Exclusive

	for k > 0 {
		p[s2] = p[s1]
		s1++
		s2--
		k--
	}

	for i := s1; i < n; i++ {
		fmt.Printf("%d ", p[i])
	}
	for i := s2 + 1; i < 2*n; i++ {
		fmt.Printf("%d ", p[i])
	}
	fmt.Printf("\n")
}
