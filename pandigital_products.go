package main

import (
	"fmt"
)

func main() {
	product_map := make(map[int]bool)

	for multiplicand := 1; multiplicand <= 100; multiplicand++ {
		dmc := digits(multiplicand)
		for multiplier := multiplicand + 1; ; multiplier++ {
			product := multiplicand * multiplier
			dmr := digits(multiplier)
			dmp := digits(product)
			if dmc+dmr+dmp > 9 {
				break
			} else if dmc+dmr+dmp == 9 {
				if candigits([]int{multiplicand, multiplier, product}) {
					product_map[product] = true
					fmt.Printf("%d * %d = %d\n", multiplicand, multiplier, product)
				}
			}
		}
	}

	sum := 0
	for k, _ := range product_map {
		sum += k
	}
	fmt.Println(sum)
}

func digits(n int) int {
	d := 0
	for n > 0 {
		d++
		n /= 10
	}
	return d
}

func candigits(tri []int) bool {
	digits := make([]bool, 9)
	for _, d := range tri {
		for d > 0 {
			i := d%10 - 1
			d /= 10
			if i < 0 {
				return false
			} else if digits[i] {
				return false
			} else {
				digits[i] = true
			}
		}
	}
	return true
}
