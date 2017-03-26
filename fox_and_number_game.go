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

	array := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		array[i] = a
	}

	has_inequality := true
	for has_inequality {
		has_inequality = false
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if array[i] != array[j] {
					has_inequality = true
					c := sub_algorithm(array[i], array[j])
					array[i] = c
					array[j] = c
				}
			}
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += array[i]
	}
	fmt.Println(sum)
}

func sub_algorithm(a, b int) int {
	if a == b {
		return a
	} else if b > a {
		a, b = b, a
	}
	return sub_algorithm(a-b, b)
}
