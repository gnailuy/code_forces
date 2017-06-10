package main

import (
	"fmt"
	"math"
)

func main() {
	var l int
	if _, err := fmt.Scanf("%d\n", &l); err != nil {
		return
	}

	min := int(math.Pow(float64(10), float64(l-1)))
	max := int(math.Pow(float64(10), float64(l))) - 1

	store := make(map[[2]int]bool)

	for numerator := min; numerator <= max; numerator++ {
		for denominator := numerator + min; denominator <= max; denominator++ {
			if numerator%10 == 0 && denominator%10 == 0 {
				continue
			}
			cn, cd := cancel_digits(numerator, denominator)
			if cn == numerator && cd == denominator {
				continue
			}
			if cn <= 0 || cd <= 0 {
				continue
			}
			g := gcd(numerator, denominator)
			n := numerator / g
			d := denominator / g
			cg := gcd(cn, cd)
			cn /= cg
			cd /= cg
			if cn == n && cd == d {
				store[[2]int{numerator, denominator}] = true
			}
		}
	}

	n_new := 1
	d_new := 1
	for k, _ := range store {
		n_new *= k[0]
		d_new *= k[1]
	}
	g := gcd(n_new, d_new)

	fmt.Println(n_new/g, d_new/g)
}

func cancel_digits(n, d int) (int, int) {
	n_digs := to_digits(n)
	d_digs := to_digits(d)
	cancelled := true
	for cancelled {
		cancelled = false
		for i := 0; i < len(n_digs); i++ {
			for j := 0; j < len(d_digs); j++ {
				if n_digs[i] == d_digs[j] {
					n_digs = append(n_digs[:i], n_digs[(i+1):]...)
					d_digs = append(d_digs[:j], d_digs[(j+1):]...)
					cancelled = true
					break
				}
			}
			if cancelled {
				break
			}
		}
	}
	return to_value(n_digs), to_value(d_digs)
}

func to_digits(n int) []int {
	var d []int
	for n > 0 {
		i := n % 10
		n /= 10
		d = append(d, i)
	}
	return d
}

func to_value(d []int) int {
	n := 0
	for i := 0; i < len(d); i++ {
		n += d[i] * int(math.Pow(float64(10), float64(i)))
	}
	return n
}

func gcd(a, b int) int {
	if a == b {
		return a
	} else if a > b {
		return gcd(b, a)
	} else if a == 0 {
		return b
	} else {
		r := b % a
		return gcd(r, a)
	}
}
