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
	b := make([]int, n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		a[i] = x
	}
	for i := 0; i < n && scanner.Scan(); i++ {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		b[i] = x
	}

	max := -1
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sum := f(a, i, j) + f(b, i, j)
			if sum > max {
				max = sum
			}
		}
	}
	fmt.Println(max)
}

func f(a []int, i, j int) int {
	or := 0
	for k := i; k <= j; k++ {
		or |= a[k]
	}
	return or
}
