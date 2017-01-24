package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// Read n and d
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}
	scanner.Scan()
	d, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	// Read friends
	money := make([]int, n)
	friendship := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		m, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		money[i] = m

		scanner.Scan()
		s, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		friendship[i] = s
	}

	if n == 1 {
		fmt.Println(friendship[0])
		return
	}

	quick_sort(money, friendship, 0, n-1)

	var max_ffactor int64 = 0
	var ffactor int64 = int64(friendship[0])

	start := 0
	end := 1
	for end < n && money[end] < money[start]+d {
		ffactor += int64(friendship[end])
		end++
	}
	max_ffactor = ffactor
	for end < n {
		// Move start pointer forward
		p := start
		for p < n && money[p] == money[start] {
			ffactor -= int64(friendship[p])
			p++
		}
		start = p
		for end < n && money[end] < money[start]+d {
			ffactor += int64(friendship[end])
			end++
		}
		if ffactor > max_ffactor {
			max_ffactor = ffactor
		}
	}
	if ffactor > max_ffactor {
		max_ffactor = ffactor
	}

	fmt.Println(max_ffactor)
}

func quick_sort(array, company []int, start, end int) {
	if start < end {
		p := partition(array, company, start, end)
		quick_sort(array, company, start, p-1)
		quick_sort(array, company, p+1, end)
	}
}

func partition(array, company []int, start, end int) int {
	pivot := array[end]
	i := start
	for j := start; j < end; j++ {
		if array[j] <= pivot {
			swap(array, i, j)
			swap(company, i, j)
			i++
		}
	}
	swap(array, i, end)
	swap(company, i, end)
	return i
}

func swap(array []int, i, j int) {
	if i != j {
		tmp := array[i]
		array[i] = array[j]
		array[j] = tmp
	}
}
