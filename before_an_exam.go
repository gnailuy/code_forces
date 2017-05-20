package main

import (
	"fmt"
)

func main() {
	var d, sum int
	if _, err := fmt.Scanf("%d %d\n", &d, &sum); err != nil {
		return
	}

	ins := make([][2]int, d) // Instructions

	for i := 0; i < d; i++ {
		if _, err := fmt.Scanf("%d %d\n", &ins[i][0], &ins[i][1]); err != nil {
			return
		}
	}

	res := make([]int, d)
	for i := 0; i < d; i++ {
		placed := false
		for di := ins[i][0]; di <= ins[i][1]; di++ {
			remain := sum - di
			min := 0
			max := 0
			for j := i + 1; j < d; j++ {
				min += ins[j][0]
				max += ins[j][1]
			}
			if remain >= min && remain <= max {
				res[i] = di
				sum = remain
				placed = true
				break
			}
		}
		if !placed {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
	for i := 0; i < d; i++ {
		fmt.Printf("%d ", res[i])
	}
	fmt.Println()
}
