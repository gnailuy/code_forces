package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var m, n int
	if _, err := fmt.Scanf("%d %d\n", &m, &n); err != nil {
		return
	}

	B := make([][]int, m)
	for i := 0; i < m; i++ {
		B[i] = make([]int, n)
	}

	zero_rows := make(map[int]bool)
	zero_cols := make(map[int]bool)

	has_one := false

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < m; i++ {
		for j := 0; j < n && scanner.Scan(); j++ {
			a, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return
			}
			B[i][j] = a
			if B[i][j] == 0 {
				zero_rows[i] = true
				zero_cols[j] = true
			} else if !has_one {
				has_one = true
			}
		}
	}

	if !has_one || m == 1 || n == 1 {
		if !has_one || len(zero_rows) == 0 {
			fmt.Println("YES")
			for i := 0; i < m; i++ {
				for j := 0; j < n; j++ {
					fmt.Printf("%d ", B[i][j])
				}
				fmt.Println()
			}
		} else {
			fmt.Println("NO")
		}
		return
	}

	impossible := false
	if len(zero_rows) >= m || len(zero_cols) >= n {
		impossible = true
	} else {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if B[i][j] == 1 {
					if zero_rows[i] && zero_cols[j] {
						impossible = true
						break
					}
				}
			}
			if impossible {
				break
			}
		}
	}

	if impossible {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if zero_rows[i] || zero_cols[j] {
					fmt.Printf("0 ")
				} else {
					fmt.Printf("1 ")
				}
			}
			fmt.Println()
		}
	}
}
