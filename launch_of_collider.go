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
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	var directions string
	coordinates := make([]int, n)
	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)

	if _d, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		directions = re.ReplaceAllString(_d, "")
	}

	if _c, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		_c := re.ReplaceAllString(_c, "")
		ss := strings.Split(_c, " ")
		for i, s := range ss {
			x, err := strconv.Atoi(s)
			if err != nil {
				return
			}
			coordinates[i] = x
		}
	}

	min_distance := -1
	for i := 0; i < n-1; i++ {
		if directions[i] == 'R' && directions[i+1] == 'L' {
			dist := (coordinates[i+1] - coordinates[i]) / 2
			if dist < min_distance || min_distance < 0 {
				min_distance = dist
			}
		}
	}

	fmt.Println(min_distance)
}
