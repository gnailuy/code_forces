package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)

	if s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		s = re.ReplaceAllString(s, "")

		directions := map[rune]int{}
		for _, r := range s {
			directions[r] += 1
		}

		xd := abs(directions['R'] - directions['L'])
		yd := abs(directions['U'] - directions['D'])

		if (xd+yd)%2 != 0 {
			fmt.Println(-1)
		} else {
			fmt.Println((xd + yd) / 2)
		}
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}
