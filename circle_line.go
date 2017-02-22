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

	total := 0
	dist := make([]int, n)
	st := make([]int, 2)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		if i < n {
			dist[i] = a
			total += a
		} else {
			st[i-n] = a
		}
	}

	if st[0] > st[1] {
		st[0], st[1] = st[1], st[0]
	}

	path := 0
	for i := st[0] - 1; i < st[1]-1; i++ {
		path += dist[i]
	}

	if total-path < path {
		fmt.Println(total - path)
	} else {
		fmt.Println(path)
	}
}
