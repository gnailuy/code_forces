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
	var n, m, k int
	if _, err := fmt.Scanf("%d %d %d\n", &n, &m, &k); err != nil {
		return
	}

	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, m)
	}
	lost := false
	lost_at := -1

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)

	for i := 0; i < k; i++ {
		if s, err := reader.ReadString('\n'); err != nil {
			if err != io.EOF {
				return
			}
		} else if lost {
			continue
		} else {
			s = re.ReplaceAllString(s, "")
			xy := strings.Split(s, " ")
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
			board[x-1][y-1] = 1
			if lose(board, n, m, x-1, y-1) {
				lost = true
				lost_at = i + 1
			}
		}
	}

	if lost {
		fmt.Println(lost_at)
	} else {
		fmt.Println(0)
	}
}

func lose(b [][]int, n, m, x, y int) bool {
	if x > 0 && y > 0 {
		if b[x-1][y-1] == 1 && b[x-1][y] == 1 && b[x][y-1] == 1 {
			return true
		}
	}
	if x > 0 && y < m-1 {
		if b[x-1][y+1] == 1 && b[x-1][y] == 1 && b[x][y+1] == 1 {
			return true
		}
	}
	if x < n-1 && y < m-1 {
		if b[x+1][y+1] == 1 && b[x+1][y] == 1 && b[x][y+1] == 1 {
			return true
		}
	}
	if x < n-1 && y > 0 {
		if b[x+1][y-1] == 1 && b[x+1][y] == 1 && b[x][y-1] == 1 {
			return true
		}
	}
	return false
}
