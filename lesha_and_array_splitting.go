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

	sum := 0
	has_non_zero := false

	array := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		array[i] = a
		sum += a
		if a != 0 {
			has_non_zero = true
		}
	}

	if !has_non_zero {
		fmt.Println("NO")
	} else if sum != 0 {
		fmt.Println("YES")
		fmt.Println("1\n1", n)
	} else {
		s := 0
		bound := -1
		for i := 0; i < n; i++ {
			s += array[i]
			if s != 0 {
				bound = i
				break
			}
		}
		fmt.Println("YES")
		fmt.Println(2)
		fmt.Println(1, bound+1)
		fmt.Println(bound+2, n)
	}
}
