package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	var n, m, k int
	if _, err := fmt.Scanf("%d %d %d\n", &n, &m, &k); err != nil {
		return
	}

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)

	maze := make([][]rune, n+2)
	for i := 0; i < n+2; i++ {
		maze[i] = make([]rune, m+2)
	}

	has_start := false
	start := [2]int{0, 0}
	x_count := 0

	for i := 1; i <= n; i++ {
		if s, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else {
			s = re.ReplaceAllString(s, "")
			maze[i][0] = '#'
			for j := 1; j <= m; j++ {
				maze[i][j] = rune(s[j-1])
				if maze[i][j] == '.' {
					if !has_start {
						start[0] = i
						start[1] = j
						has_start = true
					} else {
						maze[i][j] = 'X'
						x_count++
					}
				}
			}
			maze[i][m+1] = '#'
		}
	}
	for i := 0; i < m+2; i++ {
		maze[0][i] = '#'
		maze[n+1][i] = '#'
	}

	starts := make(map[[2]int]bool)
	starts[start] = true

	remain := x_count - k
	for remain > 0 {
		new_starts := make(map[[2]int]bool)
		for s := range starts {
			i := s[0] - 1
			j := s[1]
			if remove_x(maze, i, j, new_starts) {
				remain--
				if remain == 0 {
					break
				}
			}

			i = s[0] + 1
			j = s[1]
			if remove_x(maze, i, j, new_starts) {
				remain--
				if remain == 0 {
					break
				}
			}

			i = s[0]
			j = s[1] - 1
			if remove_x(maze, i, j, new_starts) {
				remain--
				if remain == 0 {
					break
				}
			}

			i = s[0]
			j = s[1] + 1
			if remove_x(maze, i, j, new_starts) {
				remain--
				if remain == 0 {
					break
				}
			}
		}
		starts = new_starts
	}

	for i := 1; i <= n; i++ {
		fmt.Println(string(maze[i][1 : m+1]))
	}
}

func remove_x(maze [][]rune, i, j int, new_starts map[[2]int]bool) bool {
	if maze[i][j] == 'X' {
		maze[i][j] = '.'
		new_starts[[2]int{i, j}] = true
		return true
	}
	return false
}
