package main

import (
	"fmt"
)

func main() {
	var n, k int
	if _, err := fmt.Scanf("%d %d\n", &n, &k); err != nil {
		return
	}

	if n < k {
		fmt.Println(k)
	} else {
		if n/k >= 10000 {
			for i := n + 1; ; i++ {
				if i%k == 0 {
					fmt.Println(i)
					break
				}
			}
		} else {
			for i := 2; ; i++ {
				if i*k > n {
					fmt.Println(i * k)
					break
				}
			}
		}
	}
}
