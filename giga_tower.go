package main

import (
	"fmt"
)

func main() {
	var n int64
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	for b := 1; ; b++ {
		n++
		if is_lucky(n) {
			fmt.Println(b)
			break
		}
	}
}

func is_lucky(n int64) bool {
	var lucky bool = false
	for n != 0 {
		d := n % 10
		if d == 8 || d == -8 {
			lucky = true
			break
		}
		n /= 10
	}
	return lucky
}
