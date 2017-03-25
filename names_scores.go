package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filename := os.Args[1]
	namestr, err := ioutil.ReadFile(filename)
	check(err)

	re := regexp.MustCompile(`"|\r?\n`)
	in := re.ReplaceAllString(string(namestr), "")
	names := strings.Split(in, ",")
	merge_sort(names, 0, len(names)-1)

	score := int64(0)
	for i := 0; i < len(names); i++ {
		s_name := name_score(names[i])
		fmt.Println(i+1, names[i], s_name)
		score += int64(i+1) * s_name
	}

	fmt.Println(score)
}

func name_score(name string) int64 {
	s := int64(0)
	for i := 0; i < len(name); i++ {
		s += int64(name[i] - 'A' + 1)
	}
	return s
}

func merge(buffer []string, start, middle, end int) {
	sorted := make([]string, end-start+1)

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

func merge_sort(buffer []string, start, end int) {
	if end > start {
		middle := (end + start) / 2
		merge_sort(buffer, start, middle)
		merge_sort(buffer, middle+1, end)
		merge(buffer, start, middle, end)
	}
}
