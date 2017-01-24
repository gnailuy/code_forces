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

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	array := make([]int, n)
	for i := 0; scanner.Scan(); i++ {
		m, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		array[i] = m
	}

	quick_sort(array, 0, n-1)

	for i := 0; i < n; i++ {
		fmt.Println(array[i])
	}
}

func quick_sort(array []int, start, end int) {
	if start < end {
		p := partition(array, start, end)
		quick_sort(array, start, p-1)
		quick_sort(array, p+1, end)
	}
}

func partition(array []int, start, end int) int {
	pivot := array[end]
	i := start
	for j := start; j < end; j++ {
		if array[j] <= pivot {
			swap(array, i, j)
			i++
		}
	}
	swap(array, i, end)
	return i
}

func swap(array []int, i, j int) {
	if i != j {
		tmp := array[i]
		array[i] = array[j]
		array[j] = tmp
	}
}
