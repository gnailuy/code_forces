package main

import (
	"fmt"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	grids := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		grids[i] = make([]int, n+1)
		grids[i][0] = 1
	}
	for i := 0; i <= n; i++ {
		grids[0][i] = 1
	}

	for i := 2; i <= 2*n; i++ {
		for j := 1; j <= i/2; j++ {
			k := i - j
			if k > 0 && k <= n {
				grids[j][k] = grids[j-1][k] + grids[j][k-1]
				grids[k][j] = grids[j][k]
			}
		}
	}

	fmt.Println(grids[n][n])
}
