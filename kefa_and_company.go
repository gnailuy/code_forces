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

	merge_sort(money, friendship, 0, n-1)

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
		// Move end pointer forward
		for end < n && money[end] < money[start]+d {
			ffactor += int64(friendship[end])
			end++
		}
		// See if ffactor gets larger
		if ffactor > max_ffactor {
			max_ffactor = ffactor
		}
	}

	fmt.Println(max_ffactor)
}

func merge_sort(buffer, company []int, start, end int) {
	if end > start {
		middle := (end + start) / 2
		merge_sort(buffer, company, start, middle)
		merge_sort(buffer, company, middle+1, end)
		merge(buffer, company, start, middle, end)
	}
}

func merge(buffer, company []int, start, middle, end int) {
	sorted := make([]int, end-start+1)
	c := make([]int, end-start+1)

	i := start
	j := middle + 1
	k := 0
	for i <= middle && j <= end {
		if buffer[i] <= buffer[j] {
			sorted[k] = buffer[i]
			c[k] = company[i]
			i++
		} else {
			sorted[k] = buffer[j]
			c[k] = company[j]
			j++
		}
		k++
	}
	for i <= middle {
		sorted[k] = buffer[i]
		c[k] = company[i]
		k++
		i++
	}
	for j <= end {
		sorted[k] = buffer[j]
		c[k] = company[j]
		k++
		j++
	}

	i = start
	k = 0
	for i <= end {
		buffer[i] = sorted[k]
		company[i] = c[k]
		i++
		k++
	}
}
