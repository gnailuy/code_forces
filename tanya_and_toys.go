package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, m int
	if _, err := fmt.Scanf("%d %d\n", &n, &m); err != nil {
		return
	}

	array := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		array[i] = a
	}

	merge_sort(array, 0, n-1)

	k := 0
	var buy []int

	j := 0
	ok := true
	start := 1
	end := array[j]
	for m > 0 {
		for i := start; i < end; i++ {
			if i <= m {
				m -= i
				k++
				buy = append(buy, i)
			} else {
				ok = false
				break
			}
		}
		if !ok {
			break
		}
		j++
		start = end + 1
		if j < n {
			end = array[j]
		} else {
			end = 1000000001
		}
	}
	fmt.Println(k)
	fmt.Println(strings.Trim(fmt.Sprint(buy), "[]"))
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
