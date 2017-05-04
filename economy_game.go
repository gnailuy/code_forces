package main

import (
	"fmt"
)

func main() {
	const house int = 1234567
	const car int = 123456
	const computer int = 1234

	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	ok := false
	for i := 0; i <= n/house; i++ {
		for j := 0; j <= (n-house*i)/car; j++ {
			if (n-house*i-car*j)%computer == 0 {
				ok = true
				break
			}
		}
		if ok {
			break
		}
	}

	if ok {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
