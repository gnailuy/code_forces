package main

import (
	"fmt"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	var queue []int
	var length int = 0

	for i := 1; i <= n; i++ {
		if len(queue) > length {
			queue[length] = 1
		} else {
			queue = append(queue, 1)
		}
		length++
		for j := length - 1; j > 0; j-- {
			if queue[j] == queue[j-1] {
				queue[j-1] = queue[j-1] + 1
				length--
			} else {
				break
			}
		}
	}

	for i := 0; i < length; i++ {
		fmt.Printf("%d ", queue[i])
	}
	fmt.Println()
}
