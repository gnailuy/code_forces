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

	pair_count := make([]int, 4) // ee, eo, oe, oo
	sum := make([]int, 2)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		t, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		sum[0] += t

		scanner.Scan()
		b, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		sum[1] += b

		if t%2 == 0 {
			if b%2 == 0 {
				pair_count[0]++
			} else {
				pair_count[1]++
			}
		} else {
			if b%2 == 0 {
				pair_count[2]++
			} else {
				pair_count[3]++
			}
		}
	}

	et := sum[0]%2 == 0
	eb := sum[1]%2 == 0

	if et && eb {
		fmt.Println(0)
	} else if et && !eb || eb && !et {
		fmt.Println(-1)
	} else {
		if pair_count[1] > 0 || pair_count[2] > 0 {
			fmt.Println(1)
		} else {
			fmt.Println(-1)
		}
	}
}
