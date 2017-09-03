package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var max int
	if _, err := fmt.Scanf("%d\n", &max); err != nil {
		return
	}

	not_prime := make([]bool, max)
	not_prime[1] = true
	not_prime[2] = false
	not_prime[3] = false
	for i := 2; i < max; i++ {
		for j := i; j < max; j++ {
			if i*j < max {
				not_prime[i*j] = true
			} else {
				break
			}
		}
	}

	for i := max - 1; i > 0; i-- {
		if !not_prime[i] && is_pandigital(strconv.Itoa(i)) {
			fmt.Println(i)
			break
		}
	}
}

func is_pandigital(s string) bool {
	n := len(s)
	if n > 9 {
		return false
	} else if n <= 1 {
		return true
	}
	d_list := make([]bool, n)
	for i := 0; i < n; i++ {
		d, err := strconv.Atoi(string(s[i]))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error converting to int:", string(s[i]))
			return false
		}
		if d == 0 || d > n {
			return false
		}
		if d_list[d-1] {
			return false
		}
		d_list[d-1] = true
	}
	return true
}
