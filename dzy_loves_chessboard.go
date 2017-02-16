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

	board := make([]string, n)

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
			board[i] = re.ReplaceAllString(s, "")
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == '-' {
				fmt.Printf("-")
			} else { // board[i][j] == '.'
				if (i+j)%2 == 0 {
					fmt.Printf("B")
				} else {
					fmt.Printf("W")
				}
			}
		}
		fmt.Printf("\n")
	}
}
