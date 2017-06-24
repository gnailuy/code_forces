package main

import (
	"fmt"
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

	cnt := 0
	for i := 1; i < max; i++ {
		if !not_prime[i] {
			ok := true
			j := i
			for true {
				j = circle_right(j)
				if i == j {
					break
				}
				if not_prime[j] {
					ok = false
					break
				}
			}
			if ok {
				fmt.Println(i)
				cnt += 1
			}
		}
	}

	fmt.Println(cnt)
}

func circle_right(j int) int {
	right_most := j % 10
	left := j / 10
	for j > 0 {
		j /= 10
		right_most *= 10
	}
	return right_most/10 + left
}
