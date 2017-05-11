package main

import (
	"fmt"
)

func main() {
	var xy [3][2]int

	for i := 0; i < 3; i++ {
		if _, err := fmt.Scanf("%d %d\n", &xy[i][0], &xy[i][1]); err != nil {
			return
		}
	}

	fmt.Println(3)
	for i := 0; i < 3; i++ {
		var c, o [2]int
		for j := 0; j < 3; j++ {
			if j != i {
				for k := 0; k < 2; k++ {
					c[k] += xy[j][k]
				}
			}
		}
		for k := 0; k < 2; k++ {
			o[k] = c[k] - xy[i][k]
		}

		fmt.Println(o[0], o[1])
	}
}
