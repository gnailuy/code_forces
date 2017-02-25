package main

import (
	"fmt"
)

func main() {
	var d, L, v1, v2 int
	if _, err := fmt.Scanf("%d %d %d %d\n", &d, &L, &v1, &v2); err != nil {
		return
	}

	time := float64(L-d) / float64(v1+v2)

	fmt.Printf("%.20f\n", time)
}
