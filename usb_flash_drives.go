package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, m int
	if _, err := fmt.Scanf("%d\n%d\n", &n, &m); err != nil {
		return
	}

	u := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		u[i] = a
	}

	bubble_sort(u, n)

	disks := 0
	capacity := 0
	for i := 0; i < n; i++ {
		disks += 1
		capacity += u[i]
		if capacity >= m {
			break
		}
	}
	fmt.Println(disks)
}

func swap(a, b int, array []int) {
	tmp := array[a]
	array[a] = array[b]
	array[b] = tmp
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
