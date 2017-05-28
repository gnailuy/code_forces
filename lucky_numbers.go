package main

import (
	"fmt"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	doors := int64(0)
	for i := 1; i <= n; i++ {
		doors += (2 << uint(i-1))
	}

	fmt.Println(doors)
}
