package main

import (
	"fmt"
)

func main() {
	var l1, r1, l2, r2, k int64
	if _, err := fmt.Scanf("%d %d %d %d %d\n", &l1, &r1, &l2, &r2, &k); err != nil {
		return
	}

	if r2 < l1 || r1 < l2 {
		fmt.Println("0")
	} else {
		l := l1
		if l < l2 {
			l = l2
		}
		r := r1
		if r > r2 {
			r = r2
		}

		var prink int64 = 0
		if l <= k && k <= r {
			prink = 1
		}

		fmt.Println(r - l - prink + 1)
	}
}
