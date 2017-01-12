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

	var last_k_height int = 0
	var min_k_height int = 0
	var min_k_index int = 1

	fences := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan(); i++ {
		h, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		fences[i] = h
		if i < k {
			last_k_height += fences[i]
			min_k_height += fences[i]
		} else {
			height := last_k_height - fences[i-k] + fences[i]
			last_k_height = height
			if height < min_k_height {
				min_k_height = height
				min_k_index = i - k + 2
			}
		}
	}

	fmt.Println(min_k_index)
}
