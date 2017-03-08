package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	var n, m int
	if _, err := fmt.Scanf("%d %d\n", &n, &m); err != nil {
		return
	}

	success := 0
	bests := make([]int, m)
	gradebook := make([][]int, n)

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)
	for i := 0; i < n; i++ {
		if s, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else {
			s = re.ReplaceAllString(s, "")

			gradebook[i] = make([]int, m)
			for j := 0; j < m; j++ {
				g := int(s[j] - '0')
				gradebook[i][j] = g
				if bests[j] < g {
					bests[j] = g
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if gradebook[i][j] == bests[j] {
				success++
				break
			}
		}
	}

	fmt.Println(success)
}
