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

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)

	diagonals := make(map[[2]int]int)

	for i := 0; i < n; i++ {
		if _in, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else {
			xy := strings.Split(re.ReplaceAllString(_in, ""), " ")
			if len(xy) != 2 {
				return
			}
			x, err := strconv.Atoi(xy[0])
			if err != nil {
				return
			}
			y, err := strconv.Atoi(xy[1])
			if err != nil {
				return
			}

			bp := y - x
			bn := y + x
			diagonals[[2]int{1, bp}] += 1
			diagonals[[2]int{-1, bn}] += 1
		}
	}

	res := 0
	for _, c := range diagonals {
		if c > 1 {
			res += c * (c - 1) / 2
		}
	}
	fmt.Println(res)
}
