package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	const mmd string = "Bulbasaur"
	var s string

	re := regexp.MustCompile(`\r?\n`)

	reader := bufio.NewReader(os.Stdin)
	if _s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		s = re.ReplaceAllString(_s, "")
	}

	idx := -1
	storage := make([][]int, len(mmd))
	for i := 0; i < len(mmd); i++ {
		met := false
		for j := 0; j < idx; j++ {
			if storage[j][0] == int(mmd[i]) {
				storage[j][1]++
				met = true
				break
			}
		}
		if !met {
			idx++
			storage[idx] = make([]int, 3)
			storage[idx][0] = int(mmd[i])
			storage[idx][1] = 1
		}
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j <= idx; j++ {
			if storage[j][0] == int(s[i]) {
				storage[j][2]++
				break
			}
		}
	}

	var count int = -1
	for i := 0; i <= idx; i++ {
		c := storage[i][2] / storage[i][1]
		if c < count || count < 0 {
			count = c
		}
	}

	fmt.Println(count)
}
