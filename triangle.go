package main

import (
	"fmt"
)

func main() {
	var l [4]int
	if _, err := fmt.Scanf("%d %d %d %d\n", &l[0], &l[1], &l[2], &l[3]); err != nil {
		return
	}

	segment := false

	var t [3]int
	for i := 0; i < 4; i++ {
		k := 0
		for j := 0; j < 4; j++ {
			if j != i {
				t[k] = l[j]
				k++
			}
		}
		s := triangle(t)
		if s == 1 { // triangle
			fmt.Println("TRIANGLE")
			return
		} else if s == 2 { // Degenerate triangle
			segment = true
		}
	}

	if segment {
		fmt.Println("SEGMENT")
	} else {
		fmt.Println("IMPOSSIBLE")
	}
}

func triangle(t [3]int) int {
	sum := 0
	max := 0
	for i := 0; i < 3; i++ {
		sum += t[i]
		if t[i] > max {
			max = t[i]
		}
	}

	if sum-max == max {
		return 2
	} else if sum-max > max {
		return 1
	} else {
		return 0
	}
}
