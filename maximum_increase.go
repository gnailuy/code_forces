package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	var prev int = 0
	var inc_len int = 1
	var max_inc_len int = 0

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}

		if a > prev && prev > 0 {
			inc_len += 1
		} else {
			if max_inc_len < inc_len {
				max_inc_len = inc_len
			}
			inc_len = 1
		}
		prev = a
	}
	if max_inc_len < inc_len {
		max_inc_len = inc_len
	}

	fmt.Println(max_inc_len)
}
