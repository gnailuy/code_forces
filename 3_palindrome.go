package main

import (
	"fmt"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	if n%2 != 0 {
		fmt.Printf("a")
	}

	for i := 1; i <= n/2; i++ {
		if i%2 == 0 {
			fmt.Printf("aa")
		} else {
			fmt.Printf("bb")
		}
	}
	fmt.Println()
}
