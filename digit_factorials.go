package main

import (
	"fmt"
	"math"
)

var d_fact [10]int

func init_fact(f *[10]int) {
	f[0] = 1
	for i := 1; i < 10; i++ {
		f[i] = f[i-1] * i
	}
	fmt.Println("Initliazed fact array", *f)
}

func fact_sum(i int) int {
	res := 0
	for i > 0 {
		d := i % 10
		res += d_fact[d]
		i /= 10
	}
	return res
}

func main() {
	init_fact(&d_fact)
	sum := 0

	for i := 2; ; i++ { // Number of digits
		start := int(math.Pow(float64(10), float64(i-1))) // int64
		end := int(math.Pow(float64(10), float64(i)))     // int64
		fact_end := d_fact[9] * i

		if start >= fact_end {
			break
		}

		for j := start; j <= end; j++ {
			f := fact_sum(j)
			if f == j {
				sum += j
				fmt.Println(j)
			}
		}
	}

	fmt.Println(sum)
}
