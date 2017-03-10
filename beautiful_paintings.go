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

	paintings := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan() && i < n; i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		paintings[i] = a
	}

	merge_sort(paintings, 0, n-1)

	happiness := 0
	finished := false
	for !finished {
		finished = true

		last := -1
		for i := 0; i < n; i++ {
			if paintings[i] != -1 {
				finished = false

				if last < 0 {
					last = paintings[i]
					paintings[i] = -1
				} else if last > 0 && paintings[i] > last {
					happiness++
					last = paintings[i]
					paintings[i] = -1
				}
			}
		}
	}

	fmt.Println(happiness)
}

func merge(buffer []int, start, middle, end int) {
	sorted := make([]int, end-start+1)

	i := start
	j := middle + 1
	k := 0
	for i <= middle && j <= end {
		if buffer[i] <= buffer[j] {
			sorted[k] = buffer[i]
			i++
		} else {
			sorted[k] = buffer[j]
			j++
		}
		k++
	}
	for i <= middle {
		sorted[k] = buffer[i]
		k++
		i++
	}
	for j <= end {
		sorted[k] = buffer[j]
		k++
		j++
	}

	i = start
	k = 0
	for i <= end {
		buffer[i] = sorted[k]
		i++
		k++
	}
}

func merge_sort(buffer []int, start, end int) {
	if end > start {
		middle := (end + start) / 2
		merge_sort(buffer, start, middle)
		merge_sort(buffer, middle+1, end)
		merge(buffer, start, middle, end)
	}
}
