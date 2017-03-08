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

	votes := make([]int, n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan() && i < n; i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		votes[i] = a
	}

	bubble_sort(votes, 1, n)

	bribing := 0
	for votes[0] <= votes[1] {
		votes[0]++
		votes[1]--
		bribing++
		if len(votes) > 2 {
			for i := 2; i < n && votes[i-1] < votes[i]; i++ {
				swap(i-1, i, votes)
			}
		}
	}
	fmt.Println(bribing)
}

func swap(a, b int, array []int) {
	array[a], array[b] = array[b], array[a]
}

func bubble_sort(array []int, start, end int) {
	for i := start; i < end; i++ {
		for j := start; j < end; j++ {
			if array[i] > array[j] {
				swap(i, j, array)
			}
		}
	}
}
