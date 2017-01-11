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

	houses := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		houses[i] = a
	}

	needs := make([]int, n)

	var max_right_floors int = houses[n-1]
	for i := n - 2; i >= 0; i-- {
		if houses[i] <= max_right_floors {
			needs[i] = max_right_floors - houses[i] + 1
		} else {
			max_right_floors = houses[i]
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", needs[i])
	}
	fmt.Println()
}
