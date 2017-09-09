package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run coded_triangle_numbers.go filename.txt")
		os.Exit(-1)
	}
	filename := os.Args[1]

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: read file error")
		os.Exit(-1)
	}

	words := strings.Split(string(bytes), ",")
	code_map := make(map[int]int)
	max_code := -1
	for _, w := range words {
		w := strings.Replace(w, "\"", "", -1)
		c := to_code(w)
		if c > max_code {
			max_code = c
		}
		code_map[c] += 1 // Default map value is zero
	}

	t_cnt := 0
	for i := 1; ; i++ {
		triangle := i * (i + 1) / 2
		if triangle > max_code {
			break
		}
		cnt, ok := code_map[triangle]
		if ok {
			t_cnt += cnt
		}
	}

	fmt.Println(t_cnt)
}

func to_code(s string) int {
	code := 0
	for _, l := range s {
		code += (int(l-'A') + 1)
	}
	return code
}
