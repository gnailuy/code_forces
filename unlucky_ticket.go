package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	a := make([]int, 2*n)

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)
	if s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		s = re.ReplaceAllString(s, "")
		for i := 0; i < 2*n; i++ {
			a[i] = int(s[i] - '0')
		}
	}

	merge_sort(a, 0, n-1)
	merge_sort(a, n, 2*n-1)

	pos := 0
	neg := 0
	for i := 0; i < n; i++ {
		if a[i] > a[n+i] {
			pos++
		} else if a[i] < a[n+i] {
			neg++
		}
	}

	if pos == n || neg == n {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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
