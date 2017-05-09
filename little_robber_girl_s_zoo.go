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

	h := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		h[i] = a
	}

	disorder := true
	for disorder {
		sp := -1
		ep := -1
		for i := 0; i < n-1; {
			if sp < 0 {
				if h[i] > h[i+1] {
					h[i], h[i+1] = h[i+1], h[i]
					sp = i
					ep = i + 1
					i += 2
				} else {
					i += 1
					continue
				}
			} else {
				if h[i] >= h[i+1] {
					h[i], h[i+1] = h[i+1], h[i]
					ep = i + 1
					i += 2
				} else {
					ep = i - 1
					break
				}
			}
		}
		if sp < 0 {
			disorder = false
		} else {
			fmt.Println(sp+1, ep+1)
		}
	}
}
