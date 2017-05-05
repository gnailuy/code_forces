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

	w := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		w[i] = a
	}

	p := 0
	for i := 0; i < n; i++ {
		p += w[i] / k
		if w[i]%k != 0 {
			p += 1
		}
	}
	if p%2 == 0 {
		fmt.Println(p / 2)
	} else {
		fmt.Println(p/2 + 1)
	}
}
