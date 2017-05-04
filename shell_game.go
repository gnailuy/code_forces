package main

import (
	"fmt"
)

func main() {
	situations := map[int][]int{
		0: []int{0, 1, 2},
		1: []int{1, 0, 2},
		2: []int{1, 2, 0},
		3: []int{2, 1, 0},
		4: []int{2, 0, 1},
		5: []int{0, 2, 1},
	}

	var n, x int
	if _, err := fmt.Scanf("%d\n%d\n", &n, &x); err != nil {
		return
	}
	n = n % 6

	fmt.Println(situations[n][x])
}
