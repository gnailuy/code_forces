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

	var max_height int = 0
	for i := 1; i <= array[0][0]; i++ {
		root := array[0][i]
		height := 1 + calculate_height(root, array)
		if height > max_height {
			max_height = height
		}
	}

	fmt.Println(max_height)
}

func calculate_height(root int, array [][]int) int {
	if array[root][0] > 0 {
		var max_height int = 0
		for i := 1; i <= array[root][0]; i++ {
			height := calculate_height(array[root][i], array)
			if height > max_height {
				max_height = height
			}
		}
		return max_height + 1
	} else {
		return 0
	}
}
