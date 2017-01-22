package main

import (
	"fmt"
)

func main() {
	var l, r int64
	if _, err := fmt.Scanf("%d %d\n", &l, &r); err != nil {
		return
	}

	if r-l < 2 || (r-l < 3 && l%2 != 0) {
		fmt.Println(-1)
	} else if l%2 == 0 {
		fmt.Println(l, l+1, l+2)
	} else {
		fmt.Println(l+1, l+2, l+3)
	}
}
