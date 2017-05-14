package main

import (
	"fmt"
)

func main() {
	var l, r int
	if _, err := fmt.Scanf("%d %d\n", &l, &r); err != nil {
		return
	}

	if l != r {
		fmt.Println(2)
	} else {
		fmt.Println(l)
	}
}
