package main

import (
	"fmt"
)

//  -0-
// |   |
// 1   2
// |   |
//  -3-
// |   |
// 4   5
// |   |
//  -6-

func main() {
	digits := [10]string{
		// Index: 0654 3210
		"01110111", // 0
		"00100100", // 1
		"01011101", // 2
		"01101101", // 3
		"00101110", // 4
		"01101011", // 5
		"01111011", // 6
		"00100101", // 7
		"01111111", // 8
		"01101111", // 9
	}

	good := make(map[int]map[int]bool)
	for i := 0; i < 10; i++ {
		good[i] = make(map[int]bool)
		for j := 0; j < 10; j++ {
			if i != j {
				if j_is_good_for_i(digits[i], digits[j]) {
					good[i][j] = true
				}
			}
		}
	}

	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	d1 := n / 10
	d2 := n - d1*10

	g1 := len(good[d1]) + 1
	g2 := len(good[d2]) + 1

	fmt.Println(g1 * g2)
}

func j_is_good_for_i(i, j string) bool {
	for k := 0; k < 8; k++ {
		if i[k] == '1' && j[k] == '0' {
			return false
		}
	}
	return true
}
