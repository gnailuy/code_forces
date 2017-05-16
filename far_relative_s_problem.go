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
	const days_of_year int = 366

	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	days := make([][2]int, days_of_year+1)

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)
	for i := 0; i < n; i++ {
		if s, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			}
			return
		} else {
			s = re.ReplaceAllString(s, "")

			idx := 0 // 0 for Female
			if s[0] == 'M' {
				idx = 1
			}

			ss := strings.Split(s, " ")
			if len(ss) != 3 {
				return
			}

			start, err := strconv.Atoi(ss[1])
			if err != nil {
				return
			}
			end, err := strconv.Atoi(ss[2])
			if err != nil {
				return
			}

			for j := start; j <= end; j++ {
				days[j][idx] += 1
			}
		}
	}

	max_half := 0
	for i := 1; i <= days_of_year; i++ {
		m := min(days[i][0], days[i][1])
		if m > max_half {
			max_half = m
		}
	}

	fmt.Println(max_half * 2)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
