package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	a := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		_a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		a[i] = _a
	}

	swap := 0
	var swaps [][2]int

	for i := n - 1; i > 0; i-- {
		max := -1
		max_idx := -1
		for j := 0; j < i; j++ {
			if max_idx < 0 || a[j] > max {
				max = a[j]
				max_idx = j
			}
		}
		if max > a[i] {
			swap++
			swaps = append(swaps, [2]int{max_idx, i})
			a[max_idx], a[i] = a[i], a[max_idx]
		}
	}

	fmt.Println(swap)
	for i := 0; i < swap; i++ {
		fmt.Println(swaps[i][0], swaps[i][1])
	}
}
