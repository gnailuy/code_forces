package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	const b_size int = 4

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)

	board := [b_size]string{}

	for i := 0; i < b_size; i++ {
		if s, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else {
			s = re.ReplaceAllString(s, "")
			board[i] = s
		}
	}

	ok := false
	candidate := make([]rune, b_size)
	n := 0

	for i := 0; i < b_size; i++ {
		// Horizontally
		for j := 0; j < b_size; j++ {
			candidate[j] = rune(board[i][j])
		}
		if is_ok(candidate, b_size) {
			ok = true
			break
		}

		// Vertically
		for j := 0; j < b_size; j++ {
			candidate[j] = rune(board[j][i])
		}
		if is_ok(candidate, b_size) {
			ok = true
			break
		}
	}

	if ok {
		fmt.Println("YES")
		return
	}

	dia_array := make([][b_size]rune, (b_size*2-1)*2)
	for j := 0; j < b_size; j++ {
		off0 := j
		off1 := (b_size*2 - 1) + (b_size - j - 1)
		for i := 0; i < b_size; i++ {
			dia_array[i+off0][j] = rune(board[i][j])
			dia_array[i+off1][j] = rune(board[i][j])
		}
	}

	for i := 0; i < (b_size*2-1)*2; i++ {
		n = 0
		for j := 0; j < b_size; j++ {
			if dia_array[i][j] != dia_array[0][b_size-1] {
				candidate[n] = dia_array[i][j]
				n++
			}
		}
		if is_ok(candidate, n) {
			ok = true
			break
		}
	}

	if ok {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func is_ok(candidate []rune, n int) bool {
	if n < 3 {
		return false
	} else {
		for e := 3; e <= n; e++ {
			xs := 0
			dots := 0
			for i := e - 3; i < e; i++ {
				if candidate[i] == 'x' {
					xs++
				} else if candidate[i] == '.' {
					dots++
				} else { // 'o'
					continue
				}
			}
			if xs == 2 && dots == 1 {
				return true
			}
		}
	}
	return false
}
