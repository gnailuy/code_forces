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

	prices := make([]int, n+1)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 1; i <= n && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		prices[i] = a
	}

	for i := 1; i < m || i < n-m+1; i++ {
		pl := 0
		pr := 0
		if m-i > 0 {
			if prices[m-i] > 0 && prices[m-i] <= k {
				pl = prices[m-i]
			}
		}
		if m+i <= n {
			if prices[m+i] > 0 && prices[m+i] <= k {
				pr = prices[m+i]
			}
		}
		if pl > 0 || pr > 0 {
			fmt.Println(i * 10)
			break
		}
	}
}
