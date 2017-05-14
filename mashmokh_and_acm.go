package main

import (
	"fmt"
)

func main() {
	const modulo int = 1000000007

	var n, k int
	if _, err := fmt.Scanf("%d %d\n", &n, &k); err != nil {
		return
	}

	cs := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		cs[i] = make([]int, k+1)
		cs[i][1] = 1
	}
	for j := 1; j <= k; j++ {
		cs[1][j] = 1
	}

	div_maps := make([][]int, n+1)
	for i := 2; i <= n; i++ {
		div_maps[i] = []int{1, i}
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				div_maps[i] = append(div_maps[i], j)
			}
		}
	}

	for i := 2; i <= n; i++ {
		for j := 2; j <= k; j++ {
			for p := 0; p < len(div_maps[i]); p++ {
				cs[i][j] += cs[div_maps[i][p]][j-1]
				cs[i][j] %= modulo
			}
		}
	}

	cnt := 0
	for i := 1; i <= n; i++ {
		cnt += cs[i][k]
		cnt %= modulo
	}
	fmt.Println(cnt)
}
