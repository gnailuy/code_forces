package main

import (
	"fmt"
)

func main() {
	var n, m int
	if _, err := fmt.Scanf("%d %d\n", &n, &m); err != nil {
		return
	}

	cnts := make([]int, 5)
	for i := 1; i <= 5 && i <= n; i++ {
		start := 5 - i
		if start == 0 {
			start = 5
		}
		end := m - m%5 - i
		for end+5 <= m {
			end += 5
		}
		if end > 0 {
			cnts[i-1] = (end-start)/5 + 1
		}
	}

	res := int64(n/5) * sum(cnts)
	for i := 1; i <= n%5; i++ {
		res += int64(cnts[i-1])
	}

	fmt.Println(res)
}

func sum(array []int) int64 {
	var res int64 = 0
	for i := 0; i < len(array); i++ {
		res += int64(array[i])
	}
	return res
}
