package main

import (
	"fmt"
)

func main() {
	var n int
	var k int64
	if _, err := fmt.Scanf("%d %d\n", &n, &k); err != nil {
		return
	}

	len := int64(1)
	center := 1
	for i := 1; i < n; i++ {
		len = len*2 + 1
		center++
	}

	for center > 0 {
		cp := (len + 1) / 2
		if k == cp {
			break
		} else if k > cp {
			k -= cp
		}
		len = (len - 1) / 2
		center--
	}

	fmt.Println(center)
}
