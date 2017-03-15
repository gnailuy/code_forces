package main

import (
	"fmt"
)

func main() {
	var n int64
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	max := int64(-1)
	for i := int64(1); i <= n/2; i++ {
		if n%i == 0 {
			j := n / i
			if i > j {
				break
			} else {
				if is_lovely(j) {
					fmt.Println(j)
					return
				}
				if is_lovely(i) && i > max {
					max = i
				}
			}
		}
	}
	if max > 0 {
		fmt.Println(max)
	} else {
		fmt.Println(1)
	}
}

func is_lovely(x int64) bool {
	for i := int64(2); i*i <= x; i++ {
		if x%(i*i) == 0 {
			return false
		}
	}
	return true
}
