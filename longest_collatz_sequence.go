package main

import (
	"fmt"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	if n < 1 {
		fmt.Println("The Input Number must not be less than 1")
		return
	}

	lengths := make([]int, 3*n+1)
	lengths[1] = 1

	for i := 2; i <= n; i++ {
		len := 0
		j := i
		for true {
			if j < 3*n && lengths[j] != 0 {
				len += lengths[j]
				break
			}
			len++
			if j%2 == 0 {
				j = j / 2
			} else {
				j = j*3 + 1
			}
		}
		lengths[i] = len
	}

	max := lengths[1]
	max_i := 1
	for i := 1; i <= n; i++ {
		if lengths[i] > max {
			max = lengths[i]
			max_i = i
		}
	}
	fmt.Println(max_i)
}
