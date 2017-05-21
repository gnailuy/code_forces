package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, l, r, x int
	if _, err := fmt.Scanf("%d %d %d %d\n", &n, &l, &r, &x); err != nil {
		return
	}

	c := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		_c, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		c[i] = _c
	}

	ways := 0

	mask := (1 << uint(n)) - 1
	for mask > 0 {
		m := mask

		cnt := 0
		sum := 0
		max := -1
		min := -1
		for i := 0; i < n; i++ {
			v := m % 2
			if v == 1 {
				cnt++
				sum += c[i]
				if c[i] > max {
					max = c[i]
				}
				if c[i] < min || min < 0 {
					min = c[i]
				}
			}
			m /= 2
		}
		if cnt >= 2 && sum >= l && sum <= r && max-min >= x {
			ways++
		}

		mask -= 1
	}

	fmt.Println(ways)
}
