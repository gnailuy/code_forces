package main

import (
	"fmt"
)

func main() {
	var a [3]int
	if _, err := fmt.Scanf("%d %d %d\n", &a[0], &a[1], &a[2]); err != nil {
		return
	}
	var x [3]int
	if _, err := fmt.Scanf("%d %d %d\n", &x[0], &x[1], &x[2]); err != nil {
		return
	}

	var d [3]int
	neg := 0
	for i := 0; i < 3; i++ {
		d[i] = a[i] - x[i]
		if d[i] < 0 {
			neg += d[i]
		}
	}

	for neg < 0 {
		ok := false
		for i := 0; i < 3; i++ {
			if d[i] >= 2 {
				neg++
				d[i] -= 2
				ok = true
				break
			}
		}
		if !ok {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
