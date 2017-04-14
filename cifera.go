package main

import (
	"fmt"
)

func main() {
	var k int
	if _, err := fmt.Scanf("%d\n", &k); err != nil {
		return
	}
	var l int
	if _, err := fmt.Scanf("%d\n", &l); err != nil {
		return
	}

	importance := 0
	k_to_i := int64(k)
	for true {
		if k_to_i > int64(l) {
			fmt.Println("NO")
			break
		} else if k_to_i == int64(l) {
			fmt.Println("YES")
			fmt.Println(importance)
			break
		}
		k_to_i *= int64(k)
		importance++
	}
}
