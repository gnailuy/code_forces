package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var max int = 0

	for i := 1; ; i++ {
		a1 := i
		a2 := i * 2
		dig_str := strconv.Itoa(a1) + strconv.Itoa(a2)
		if len(dig_str) > 9 {
			break
		}
		j := 3
		for len(dig_str) < 9 {
			dig_str += strconv.Itoa(i * j)
			j++
		}
		if len(dig_str) > 9 {
			continue
		}
		if is_pandigital(dig_str) {
			dig, err := strconv.Atoi(dig_str)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error converting dig_str to int")
			}
			if dig > max {
				max = dig
			}
		}
	}

	fmt.Println(max)
}

func is_pandigital(s string) bool {
	if len(s) != 9 {
		return false
	}
	d_map := make(map[byte]bool)
	for i := 0; i < 9; i++ {
		d := s[i]
		if int(d) == '0' {
			return false
		}
		if _, exists := d_map[d]; exists {
			return false
		}
		d_map[d] = true
	}
	return true
}
