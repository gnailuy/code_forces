package main

import (
	"fmt"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	max_n := 0
	max_len := 0

	for i := 2; i < n; i++ {
		num := 10
		numerators := make(map[int]bool)
		numerators[num] = true

		len := 0
		for true {
			len++
			remain := num % i
			if remain == 0 {
				len = 0
				break
			}
			num = remain * 10
			if numerators[num] {
				break
			} else {
				numerators[num] = true
			}
		}

		if len > max_len {
			max_len = len
			max_n = i
		}
	}

	fmt.Println(max_n)
}
