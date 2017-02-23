package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, m, k int
	if _, err := fmt.Scanf("%d %d %d\n", &n, &m, &k); err != nil {
		return
	}

	wash := 0
	bowls := m
	plates := k

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan() && i < n; i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}

		if a == 1 {
			if bowls > 0 {
				bowls--
			} else {
				wash++
			}
		} else if a == 2 {
			if plates > 0 {
				plates--
			} else if bowls > 0 {
				bowls--
			} else {
				wash++
			}
		} else {
			return
		}
	}

	fmt.Println(wash)
}
