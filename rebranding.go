package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	var n, m int
	if _, err := fmt.Scanf("%d %d\n", &n, &m); err != nil {
		return
	}

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)

	name := make(map[rune][]int)

	if s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		s = re.ReplaceAllString(s, "")
		for i := 0; i < n; i++ {
			name[rune(s[i])] = append(name[rune(s[i])], i)
		}
	}

	for i := 0; i < m; i++ {
		if s, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else {
			xy := strings.Split(re.ReplaceAllString(s, ""), " ")
			x := rune(xy[0][0])
			y := rune(xy[1][0])
			if x != y {
				name[x], name[y] = name[y], name[x]
			}
		}
	}

	cs := make([]rune, n)
	ids := make([]int, n)
	p := 0
	for c, idx := range name {
		for i := 0; i < len(idx); i++ {
			cs[p] = c
			ids[p] = idx[i]
			p++
		}
	}

	merge_sort(ids, cs, 0, n-1)
	fmt.Println(string(cs))
}

func merge_sort(buffer []int, company []rune, start, end int) {
	if end > start {
		middle := (end + start) / 2
		merge_sort(buffer, company, start, middle)
		merge_sort(buffer, company, middle+1, end)
		merge(buffer, company, start, middle, end)
	}
}

func merge(buffer []int, company []rune, start, middle, end int) {
	sorted := make([]int, end-start+1)
	c := make([]rune, end-start+1)

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
