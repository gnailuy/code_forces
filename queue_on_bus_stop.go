package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, m int
	if _, err := fmt.Scanf("%d %d\n", &n, &m); err != nil {
		return
	}

	queue := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan() && i < n; i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		queue[i] = a
	}

	buses := 0
	seats := 0
	for i := 0; i < n; i++ {
		if seats < queue[i] {
			buses++
			seats = m - queue[i] // queue[i] <= m
		} else {
			seats -= queue[i]
		}
	}

	fmt.Println(buses)
}
