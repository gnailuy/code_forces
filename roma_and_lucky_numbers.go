package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, k int
	if _, err := fmt.Scanf("%d %d\n", &n, &k); err != nil {
		return
	}

	var result int = 0

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		if no_more_than_k(a, k) {
			result++
		}
	}

	fmt.Println(result)
}

func no_more_than_k(a, k int) bool {
	cnt := 0
	for a > 0 {
		d := a % 10
		if d == 4 || d == 7 {
			cnt++
			if cnt > k {
				return false
			}
		}
		a = a / 10
	}
	return true
}
