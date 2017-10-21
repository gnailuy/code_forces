package main

import (
	"fmt"
)

func main() {
	var pow11 int = 10000000000
	var max int = 1000
	var res int = 0

	for i := 1; i <= max; i++ {
		p := i
		for j := 1; j < i; j++ {
			p *= i
			if p > pow11 {
				p %= pow11
			}
		}
		res += p
	}
	fmt.Println(res % pow11)
}
