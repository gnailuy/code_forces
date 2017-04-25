package main

import (
	"fmt"
)

func main() {
	var n, a, b, c int64
	if _, err := fmt.Scanf("%d %d %d %d\n", &n, &a, &b, &c); err != nil {
		return
	}

	need := 4 - n%4
	if need == 1 {
		fmt.Println(min(min(a, b+c), c*3))
	} else if need == 2 {
		fmt.Println(min(min(a*2, b), c*2))
	} else if need == 3 {
		fmt.Println(min(min(a*3, a+b), c))
	} else if need == 4 {
		fmt.Println(0)
	}
}

func min(x, y int64) int64 {
	if x < y {
		return x
	} else {
		return y
	}
}
