package main

import (
	"fmt"
)

func main() {
	var x, y, n int
	if _, err := fmt.Scanf("%d %d\n", &x, &y); err != nil {
		return
	}
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	an := fn(x, y, n)
	if an < 0 {
		fmt.Println((1000000007 - (-an % 1000000007)) % 1000000007)
	} else {
		fmt.Println(an % 1000000007)
	}
}

func fn(x, y, n int) int {
	if n == 1 {
		return x
	} else if n == 2 {
		return y
	} else if n == 3 {
		return y - x
	} else { // n > 3
		head := [3]int{y - x, x, y}
		pos := n % 3
		sign := 1
		if ((n-1)/3)%2 != 0 {
			sign = -1
		}
		return head[pos] * sign
	}
}
