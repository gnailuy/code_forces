package main

import (
	"fmt"
)

func main() {
	var n, m, z int
	if _, err := fmt.Scanf("%d %d %d\n", &n, &m, &z); err != nil {
		return
	}

	kill := 0
	for i := m; i <= z; i += m {
		if i%n == 0 {
			kill++
		}
	}

	fmt.Println(kill)
}
