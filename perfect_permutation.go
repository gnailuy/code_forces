package main

import (
	"fmt"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	if n%2 == 0 {
		for i := 1; i <= n; i++ {
			fmt.Printf("%d ", n-i+1)
		}
		fmt.Println()
	} else {
		fmt.Println("-1")
	}
}
