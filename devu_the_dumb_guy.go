package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, x int
	if _, err := fmt.Scanf("%d %d\n", &n, &x); err != nil {
		return
	}

	c := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		c[i] = a
	}

	merge_sort(c, 0, n-1)

	time := int64(0)
	xx := int64(x)
	for i := 0; i < n; i++ {
		time += int64(c[i]) * xx
		if xx > 1 {
			xx -= 1
		}
	}
	fmt.Println(time)
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
