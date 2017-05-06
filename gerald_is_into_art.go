package main

import (
	"fmt"
)

func main() {
	var a [3][2]int
	for i := 0; i < 3; i++ {
		if _, err := fmt.Scanf("%d %d\n", &a[i][0], &a[i][1]); err != nil {
			return
		}
	}

	m := a[0][0]
	n := a[0][1]
	for j := 0; j < 2; j++ {
		s := a[1][j]
		t := a[1][(j+1)%2]
		for k := 0; k < 2; k++ {
			u := a[2][k]
			v := a[2][(k+1)%2]

			if (s+u <= m && max(t, v) <= n) || (t+v <= n && max(s, u) <= m) {
				fmt.Println("YES")
				return
			}
		}
	}

	fmt.Println("NO")
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
