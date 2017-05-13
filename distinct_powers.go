package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	not_prime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		for j := 2; j <= n; j++ {
			if i*j <= n {
				not_prime[i*j] = true
			} else {
				break
			}
		}
	}

	av_set := make(map[string]bool)

	for x := 2; x <= n; x++ {
		for b := 2; b <= n; b++ {
			x_map := to_prime_map(not_prime, x)
			for k, v := range x_map {
				x_map[k] = v * b
			}
			av := to_prime_string(x_map, not_prime, n)
			av_set[av] = true
			fmt.Println("Added: ", av)
		}
	}

	fmt.Println(len(av_set))
}

func to_prime_map(not_prime []bool, x int) map[int]int {
	res := make(map[int]int)
	remain := x
	for i := 2; i <= x; i++ {
		if !not_prime[i] {
			if remain%i == 0 {
				cnt := 0
				for remain%i == 0 {
					remain /= i
					cnt++
				}
				res[i] = cnt
			}
		}
	}
	return res
}

func to_prime_string(prime_map map[int]int, not_prime []bool, n int) string {
	var pairs []string
	for i := 2; i <= n; i++ {
		if !not_prime[i] {
			if prime_map[i] > 0 {
				pairs = append(pairs, strconv.Itoa(i)+"^"+strconv.Itoa(prime_map[i]))
			}
		}
	}
	return strings.Join(pairs, "*")
}
