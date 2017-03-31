package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var n [2]int
	var lrs [2][2][]int

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)

	for k := 0; k < 2; k++ {
		if s, err := reader.ReadString('\n'); err != nil {
			if err != io.EOF {
				return
			}
		} else {
			n[k], err = strconv.Atoi(re.ReplaceAllString(s, ""))
			if err != nil {
				return
			}
		}

		lrs[k][0] = make([]int, n[k])
		lrs[k][1] = make([]int, n[k])
		for i := 0; i < n[k]; i++ {
			if _in, err := reader.ReadString('\n'); err != nil {
				if err == io.EOF {
					break
				} else {
					return
				}
			} else {
				lr := strings.Split(re.ReplaceAllString(_in, ""), " ")
				if len(lr) != 2 {
					return
				}
				for j := 0; j < 2; j++ {
					lrs[k][j][i], err = strconv.Atoi(re.ReplaceAllString(lr[j], ""))
					if err != nil {
						return
					}
				}
			}
		}
	}

	merge_sort(lrs[0][0], 0, n[0]-1)
	merge_sort(lrs[0][1], 0, n[0]-1)
	merge_sort(lrs[1][0], 0, n[1]-1)
	merge_sort(lrs[1][1], 0, n[1]-1)

	diff := max(lrs[1][0][n[1]-1]-lrs[0][1][0], lrs[0][0][n[0]-1]-lrs[1][1][0])
	if diff > 0 {
		fmt.Println(diff)
	} else {
		fmt.Println(0)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
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
