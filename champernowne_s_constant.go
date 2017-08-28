package main

import (
	"fmt"
	"math"
)

func main() {
	max_digit := 1000000
	last_digits := 0
	digits := 0

	d := 0
	result := 1

	for i := 1; ; i++ {
		numbers := int(math.Pow10(i) - math.Pow10(i-1))
		last_digits = digits
		digits += numbers * i

		for j := d; ; j++ {
			target := int(math.Pow10(j))
			if target > digits {
				d = j
				break
			}
			target -= last_digits
			index := target - 1
			position := 1
			start := int(math.Pow10(i - 1))
			if j > 0 {
				index = (target - 1) / j
				position = target%j + 1
			}
			td := find_i_digit(start+index, j-position+1)
			result *= td
		}

		if digits > max_digit {
			break
		}
	}

	fmt.Println(result)
}

func find_i_digit(number, b_index int) int {
	for i := 0; i < b_index; i++ {
		number /= 10
	}
	return number % 10
}
