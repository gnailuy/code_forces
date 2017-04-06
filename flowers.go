package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var t, k int
	if _, err := fmt.Scanf("%d %d\n", &t, &k); err != nil {
		return
	}

	test_cases := make([][2]int, t)
	max := 0

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < t; i++ {
		for j := 0; j < 2; j++ {
			scanner.Scan()
			a, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return
			}
			test_cases[i][j] = a
			if a > max {
				max = a
			}
		}
	}

	const modulo int64 = 1000000007

	ways := make([]int64, max+1)
	ways[0] = 1
	for i := 1; i <= max; i++ {
		if k == 1 {
			ways[i] = (ways[i-1] * 2) % modulo
		} else {
			if i < k {
				ways[i] = int64(1)
			} else {
				ways[i] = ways[i-1] + ways[i-k]
			}
		}
		if ways[i] >= modulo {
			ways[i] -= modulo
		}
	}

	ways[0] = 0
	for i := 1; i <= max; i++ {
		ways[i] += ways[i-1]
		if ways[i] >= modulo {
			ways[i] -= modulo
		}
	}

	for i := 0; i < t; i++ {
		w := ways[test_cases[i][1]] - ways[test_cases[i][0]-1]
		if w < 0 {
			w += modulo
		}
		fmt.Println(w)
	}
}
