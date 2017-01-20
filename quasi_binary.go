package main

import (
	"fmt"
)

func main() {
	const max_digits int = 7

	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	digits := make([]byte, max_digits)
	max_digit := 0
	l := 0
	m := n
	for m > 0 {
		digits[l] = byte(m % 10)
		if max_digit < int(digits[l]) {
			max_digit = int(digits[l])
		}
		m = m / 10
		l++
	}

	fmt.Println(max_digit)
	for n > 0 {
		var output bool = false
		for i := l - 1; i >= 0; i-- {
			if digits[i] > 0 {
				output = true
				fmt.Printf("1")
				digits[i] -= 1

				sub := 1
				for j := 0; j < i; j++ {
					sub *= 10
				}
				n -= sub
			} else {
				if output {
					fmt.Printf("0")
				}
			}
		}
		fmt.Printf(" ")
	}
	fmt.Println()
}
