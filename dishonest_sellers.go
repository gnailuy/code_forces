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

	ab := make([][2]int, n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for j := 0; j < 2; j++ {
		for i := 0; i < n && scanner.Scan(); i++ {
			_in, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return
			}
			ab[i][j] = _in
		}
	}

	merge_sort(ab, 0, n-1)

	money := 0
	for i := 0; i < n; i++ {
		if i < k {
			money += ab[i][0]
		} else if ab[i][0]-ab[i][1] < 0 {
			money += ab[i][0]
		} else {
			money += ab[i][1]
		}
	}
	fmt.Println(money)
}

func le(a, b [2]int) bool {
	return (a[0] - a[1]) <= (b[0] - b[1])
}

func merge(buffer [][2]int, start, middle, end int) {
	sorted := make([][2]int, end-start+1)

	i := start
	j := middle + 1
	k := 0
	for i <= middle && j <= end {
		if le(buffer[i], buffer[j]) {
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

func merge_sort(buffer [][2]int, start, end int) {
	if end > start {
		middle := (end + start) / 2
		merge_sort(buffer, start, middle)
		merge_sort(buffer, middle+1, end)
		merge(buffer, start, middle, end)
	}
}
