package main

import (
	"fmt"
	"strconv"
)

func main() {
	var max int = 1000000

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

	for i := 1; i < max; i++ {
		if !not_prime[i] {
			i_len := uint(len(strconv.Itoa(i)))
			max_mask := 1 << i_len
			for mask := 1; mask < max_mask; mask++ {
				ones := find_ones(mask)

				first_prime := -1
				num_non_prime := 0
				prime_list := []int{}

				start := 0
				if ones[int(i_len)-1] {
					start = 1
				}

				for d := start; d <= 9; d++ {
					// Replace digital at ones with d
					rep_i := replace_digital(i, int(i_len), d, ones)
					if !not_prime[rep_i] {
						if first_prime < 0 {
							first_prime = rep_i
						}
						prime_list = append(prime_list, rep_i)
					} else {
						num_non_prime += 1
					}
					if num_non_prime+start > 2 {
						break
					}
				}

				if num_non_prime+start == 2 {
					fmt.Println("Ans:", first_prime)
					fmt.Println("Primes:", prime_list)
					return
				}
			}
		}
	}
}

func find_ones(mask int) map[int]bool {
	ones := make(map[int]bool)
	i := 0
	for mask > 0 {
		if mask&1 == 1 {
			ones[i] = true
		}
		i += 1
		mask = mask >> 1
	}
	return ones
}

func replace_digital(num, num_len, d int, ones map[int]bool) int {
	for i := 0; i < num_len; i++ {
		if ones[i] {
			p := 1
			t_num := num
			for j := 0; j < i; j++ {
				p *= 10
				t_num /= 10
			}
			rd := t_num % 10
			num -= rd * p
			num += d * p
		}
	}
	return num
}
