package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; ; i++ {
		ok := true
		i_map := to_char_map(strconv.Itoa(i))
		for f := 2; f <= 6; f++ {
			j := i * f
			j_map := to_char_map(strconv.Itoa(j))
			if !equal(i_map, j_map) {
				ok = false
				break
			}
		}
		if ok {
			fmt.Println(i)
			for f := 2; f <= 6; f++ {
				fmt.Println(i * f)
			}
			break
		}
	}
}

func to_char_map(s string) map[byte]int {
	n := len(s)
	c_map := make(map[byte]int)
	for i := 0; i < n; i++ {
		c_map[s[i]] += 1
	}
	return c_map
}

func equal(a, b map[byte]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
