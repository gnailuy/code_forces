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

	sum := 0
	w := make([]int, 7)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < 7 && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		w[i] = a
		sum += a
	}

	n = n % sum
	if n == 0 {
		for i := 6; i >= 0; i-- {
			if w[i] != 0 {
				fmt.Println(i + 1)
				return
			}
		}
	}
	for i := 0; i < 7; i++ {
		n -= w[i]
		if n <= 0 {
			fmt.Println(i + 1)
			break
		}
	}
}
