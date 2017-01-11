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

	array := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		array[i] = make([]int, n+1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 1; scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		if a == -1 {
			a = 0
		}
		array[a][array[a][0]+1] = i
		array[a][0]++
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			fmt.Printf("%d ", array[i][j])
		}
		fmt.Println()
	}
	// Construct the tree and record the max height

	// fmt.Println(max_height)
}
