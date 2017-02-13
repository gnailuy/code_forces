package main

import (
	"fmt"
)

func main() {
	var n, s int
	if _, err := fmt.Scanf("%d %d\n", &n, &s); err != nil {
		return
	}

	var max_sweets int = -1
	for i := 0; i < n; i++ {
		var x, y int
		if _, err := fmt.Scanf("%d %d\n", &x, &y); err != nil {
			return
		}

		if s >= x+1 || (y == 0 && s >= x) {
			if y > 0 && 100-y > max_sweets {
				max_sweets = 100 - y
			} else if y == 0 && max_sweets < 0 {
				max_sweets = 0
			}
		}
	}

	fmt.Println(max_sweets)
}
