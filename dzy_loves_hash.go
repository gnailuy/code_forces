package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var p, n int
	if _, err := fmt.Scanf("%d %d\n", &p, &n); err != nil {
		return
	}

	store := make(map[int]bool)
	conflict := -1

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		hash := h(a, p)
		if store[hash] {
			if conflict < 0 {
				conflict = i + 1
			}
		} else {
			store[hash] = true
		}
	}

	fmt.Println(conflict)
}

func h(x, p int) int {
	return x % p
}
