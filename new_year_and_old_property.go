package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int64
	if _, err := fmt.Scanf("%d %d\n", &a, &b); err != nil {
		return
	}

	as, ac, ai := to_binary_string(a)
	bs, bc, bi := to_binary_string(b)

	var c [3]int

	if len(as) < len(bs) {
		if ac == 0 {
			c[0] = 0
		} else {
			c[0] = len(as) - ai
		}
		for i := len(as) + 1; i <= len(bs)-1; i++ {
			c[1] += i - 1
		}
		if bc == 0 {
			c[2] = len(bs) - 1
		} else {
			c[2] = bi - 1
			if bc == 1 {
				c[2] += 1
			}
		}
	} else if len(as) == len(bs) {
		c[1] = bi - ai
		if bc == 1 {
			c[2] = 1
		}
	}

	fmt.Println(c[0] + c[1] + c[2])
}

func to_binary_string(n int64) (string, int, int) {
	out := ""
	zero_count := 0
	first_zero_index := -1
	for i := 0; n > 0; i++ {
		if n%2 == 0 {
			zero_count++
			first_zero_index = i
		}
		out = strconv.FormatInt(n%2, 10) + out
		n /= 2
	}
	return out, zero_count, len(out) - first_zero_index - 1
}
