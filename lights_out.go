package main

import (
	"fmt"
)

func main() {
	const n int = 3

	var l [n][n]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			l[i][j] = 1
		}
	}

	var p [n][n]int
	for i := 0; i < n; i++ {
		if _, err := fmt.Scanf("%d %d %d\n", &p[i][0], &p[i][1], &p[i][2]); err != nil {
			return
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if p[i][j]%2 != 0 {
				l[i][j] = (l[i][j] + 1) % 2
				if i > 0 {
					l[i-1][j] = (l[i-1][j] + 1) % 2
				}
				if i < n-1 {
					l[i+1][j] = (l[i+1][j] + 1) % 2
				}
				if j > 0 {
					l[i][j-1] = (l[i][j-1] + 1) % 2
				}
				if j < n-1 {
					l[i][j+1] = (l[i][j+1] + 1) % 2
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d", l[i][j])
		}
		fmt.Println()
	}
}
