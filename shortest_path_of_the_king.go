package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	var c [2][2]int

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)
	for i := 0; i < 2; i++ {
		if s, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else {
			s = re.ReplaceAllString(s, "")
			c[i][0] = int(s[0] - 'a')
			c[i][1] = int(s[1] - '1')
		}
	}

	var steps int = 0
	var moves [8]string

	for c[0][0] != c[1][0] || c[0][1] != c[1][1] {
		x := ' '
		if c[0][0] < c[1][0] {
			c[0][0] += 1
			x = 'R'
		} else if c[0][0] > c[1][0] {
			c[0][0] -= 1
			x = 'L'
		}

		y := ' '
		if c[0][1] < c[1][1] {
			c[0][1] += 1
			y = 'U'
		} else if c[0][1] > c[1][1] {
			c[0][1] -= 1
			y = 'D'
		}
		if ' ' == x {
			moves[steps] = string(y)
		} else if ' ' == y {
			moves[steps] = string(x)
		} else {
			moves[steps] = string(x) + string(y)
		}
		steps++
	}

	fmt.Println(steps)
	for i := 0; i < steps; i++ {
		fmt.Println(moves[i])
	}
}
