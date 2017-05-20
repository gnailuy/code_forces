package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int
	if _, err := fmt.Scanf("%d %d\n%d %d\n", &a, &b, &c, &d); err != nil {
		return
	}

	offset := make(map[int]bool)

	for i := 0; ; i++ {
		t := i*a + b
		if t-d >= 0 {
			if (t-d)%c == 0 {
				fmt.Println(t)
				return
			}
			l := (t - d) / c
			off := t - (l*c + d)
			if offset[off] {
				break
			} else {
				offset[off] = true
			}
		}
	}

	fmt.Println(-1)
}
