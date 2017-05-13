package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	var c [2]int
	if _, err := fmt.Scanf("%d %d %d\n", &n, &c[0], &c[1]); err != nil {
		return
	}

	w := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		w[i] = a
	}

	merge_sort(w, 0, n-1)
	if c[0] > c[1] {
		c[0], c[1] = c[1], c[0]
	}

	var sum [2]int64
	p := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < c[i]; j++ {
			sum[i] += int64(w[p])
			p++
		}
	}

	mean := float64(sum[0])/float64(c[0]) + float64(sum[1])/float64(c[1])

	fmt.Printf("%.8f\n", mean)
}

func merge(buffer []int, start, middle, end int) {
	sorted := make([]int, end-start+1)

	i := start
	j := middle + 1
	k := 0
	for i <= middle && j <= end {
		if buffer[i] >= buffer[j] {
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
