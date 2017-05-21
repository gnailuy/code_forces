package main

import (
	"fmt"
	"strconv"
	"strings"
)

// x = b·s(x)^a + c

func main() {
	const max_x int64 = 1000000000
	const max_digits int64 = 9

	var a, b, c int64
	if _, err := fmt.Scanf("%d %d %d\n", &a, &b, &c); err != nil {
		return
	}

	cnt := 0
	var all_res []string

	for i := int64(1); i <= 9*max_digits; i++ {
		y := b*power(i, a) + c
		if y < max_x && s(y) == i {
			cnt++
			all_res = append(all_res, strconv.FormatInt(y, 10))
		}
	}

	fmt.Println(cnt)
	fmt.Println(strings.Join(all_res, " "))
}

func s(x int64) int64 {
	res := int64(0)
	for x > 0 {
		res += x % 10
		x /= 10
	}
	return res
}

func power(sx, n int64) int64 {
	res := sx
	for i := int64(1); i < n; i++ {
		res *= sx
	}
	return res
}
