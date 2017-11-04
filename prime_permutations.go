package main

import (
	"fmt"
	"strconv"
)

func main() {
	var min int = 1000
	var max int = 10000

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

	permutation_map := make(map[string][]int)

	for i := min; i < max; i++ {
		if !not_prime[i] {
			key := int_to_key(i)
			if key != "" {
				if list, ok := permutation_map[key]; ok {
					permutation_map[key] = append(list, i)
				} else {
					list = make([]int, 0, 24)
					permutation_map[key] = append(list, i)
				}
			}
		}
	}

	for _, list := range permutation_map {
		if len(list) >= 3 {
			for i := 0; i < len(list); i++ {
				for j := i + 1; j < len(list); j++ {
					for k := j + 1; k < len(list); k++ {
						if list[j]-list[i] == list[k]-list[j] {
							fmt.Println(list[i], list[j], list[k])
						}
					}
				}
			}
		}
	}
}

func int_to_key(s int) string {
	d_list := make([]int, 10)
	for s > 0 {
		d := s % 10
		d_list[d] += 1
		s /= 10
	}
	key := ""
	for i := 0; i < 10; i++ {
		for j := 0; j < d_list[i]; j++ {
			key += strconv.Itoa(i)
		}
	}
	return key
}
