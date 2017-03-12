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

	if n == 1 {
		var trash int
		if _, err := fmt.Scanf("%d\n", &trash); err != nil {
			return
		}
		fmt.Println(1)
		return
	}

	array := make([][]int64, n+1)
	array[n] = make([]int64, n+1)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	e := [2]int{-1, -1}
	diagonal := [2]int64{0, 0}
	target := int64(-1)
	met_zero := false

	for i := 0; i < n; i++ {
		array[i] = make([]int64, n+1)
		r_sum := int64(0)
		for j := 0; j < n && scanner.Scan(); j++ {
			in, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return
			}
			a := int64(in)
			array[i][j] = a
			r_sum += a
			array[n][j] += a

			if i == j {
				diagonal[0] += a
			}
			if i+j == n-1 {
				diagonal[1] += a
			}
			if a == 0 {
				e[0] = i
				e[1] = j
				met_zero = true
			}
		}
		array[i][n] = r_sum
		if target < 0 && !met_zero {
			target = r_sum
		}
		if met_zero {
			met_zero = false
		}
	}

	magic := target - array[n][e[1]]
	if target != array[e[0]][n]+magic {
		fmt.Println(-1)
		return
	}

	if e[0] != e[1] {
		if diagonal[0] != target {
			fmt.Println(-1)
			return
		}
	} else if diagonal[0]+magic != target {
		fmt.Println(-1)
		return
	}
	if e[0]+e[1] != n-1 {
		if diagonal[1] != target {
			fmt.Println(-1)
			return
		}
	} else if diagonal[1]+magic != target {
		fmt.Println(-1)
		return
	}

	check_ok := true
	for i := 0; i < n; i++ {
		if i != e[0] {
			if array[i][n] != target {
				check_ok = false
				break
			}
		}
		if i != e[1] {
			if array[n][i] != target {
				check_ok = false
				break
			}
		}
	}

	if !check_ok || magic <= 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(magic)
	}
}
