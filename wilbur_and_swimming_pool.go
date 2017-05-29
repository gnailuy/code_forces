package main

import (
	"fmt"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	xys := [2][2]int{}
	xyc := [2]int{}

	for i := 0; i < n; i++ {
		var xy [2]int
		if _, err := fmt.Scanf("%d %d\n", &xy[0], &xy[1]); err != nil {
			return
		}
		for j := 0; j < 2; j++ {
			if xyc[j] == 0 {
				xys[j][0] = xy[j]
				xyc[j] += 1
			} else if xys[j][0] != xy[j] {
				xys[j][1] = xy[j]
				xyc[j] += 1
			}
		}
	}

	if xyc[0] < 2 || xyc[1] < 2 {
		fmt.Println(-1)
	} else {
		fmt.Println(abs(xys[0][0]-xys[0][1]) * abs(xys[1][0]-xys[1][1]))
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
