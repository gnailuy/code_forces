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

	groupA := make([]int, 5)
	groupB := make([]int, 5)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		groupA[a-1]++
	}
	for i := 0; i < n && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		groupB[a-1]++
	}

	moves := 0
	for i := 0; i < 5; i++ {
		groupA[i] -= groupB[i]
		if groupA[i]%2 != 0 {
			fmt.Println(-1)
			return
		} else if groupA[i] > 0 {
			moves += groupA[i]
		}
	}
	fmt.Println(moves / 2)
}
