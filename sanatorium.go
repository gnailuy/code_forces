package main

import (
	"fmt"
)

func main() {
	d := make([]int64, 3)
	if _, err := fmt.Scanf("%d %d %d\n", &d[0], &d[1], &d[2]); err != nil {
		return
	}

	mi, mc := max_index(d)
	if mc == 3 {
		fmt.Println(0)
	} else if mc == 2 {
		fmt.Println(d[mi]*3 - 1 - d[0] - d[1] - d[2])
	} else {
		fmt.Println(d[mi]*3 - 2 - d[0] - d[1] - d[2])
	}
}

func max_index(d []int64) (int, int) {
	max := d[0]
	max_i := 0
	max_count := 1
	for i := 1; i < len(d); i++ {
		if d[i] > max {
			max_i = i
			max = d[i]
			max_count = 1
		} else if d[i] == max {
			max_count++
		}
	}
	return max_i, max_count
}
