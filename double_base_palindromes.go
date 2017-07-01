package main

import (
	"fmt"
	"strconv"
)

func main() {
	var max int
	if _, err := fmt.Scanf("%d\n", &max); err != nil {
		return
	}

	sum := 0
	for i := 1; i < max; i++ {
		ds := to_decimal_string(i)
		bs := to_binary_string(i)
		if is_palindrome(ds) && is_palindrome(bs) {
			fmt.Println(ds, bs)
			sum += i
		}
	}

	fmt.Println(sum)
}

func to_decimal_string(n int) string {
	return strconv.Itoa(n)
}

func to_binary_string(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

func is_palindrome(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}
