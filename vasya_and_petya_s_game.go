package main

import (
	"fmt"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	if n == 1 {
		fmt.Println(0)
		return
	} else if n == 2 {
		fmt.Println(1)
		fmt.Println(2)
		return
	}

	not_prime := make([]bool, n+1)
	not_prime[1] = true
	not_prime[2] = false
	not_prime[3] = false
	for i := 2; i <= n; i++ {
		for j := 2; j <= n; j++ {
			if i*j <= n {
				not_prime[i*j] = true
			} else {
				break
			}
		}
	}

	questions := 0
	is_question := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		if !not_prime[i] {
			p := i
			for p <= n {
				if !is_question[p] {
					is_question[p] = true
					questions++
				}
				p *= i
			}
		}
	}

	fmt.Println(questions)
	for i := 1; i <= n; i++ {
		if is_question[i] {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("\n")
}
