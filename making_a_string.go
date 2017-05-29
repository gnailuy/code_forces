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

	a := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		_a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		a[i] = _a
	}

	bubble_sort(a, n)

	length := int64(a[0])
	for i := 1; i < n; i++ {
		if a[i] >= a[i-1] {
			a[i] = a[i-1] - 1
		}
		if a[i] <= 0 {
			break
		}
		length += int64(a[i])
	}
	fmt.Println(length)
}

func swap(a, b int, array []int) {
	array[a], array[b] = array[b], array[a]
}

func bubble_sort(array []int, length int) {
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if array[i] > array[j] {
				swap(i, j, array)
			}
		}
	}
}
