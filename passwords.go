package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	const max_length = 100

	var n, k int
	if _, err := fmt.Scanf("%d %d\n", &n, &k); err != nil {
		return
	}

	lens := make([]int, max_length)
	var answer string

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i <= n; i++ {
		if s, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else if i < n {
			l := len(s)
			if s[l-2] == '\r' { // Tailing with '\r\n'
				l -= 2
			} else { // Tailing with '\n'
				l -= 1
			}
			lens[l-1]++
		} else {
			answer = s
		}
	}

	answer_len := len(answer)
	if answer[answer_len-2] == '\r' {
		answer_len -= 2
	} else {
		answer_len -= 1
	}

	var tries int = 0
	for i := 0; i < answer_len-1; i++ {
		tries += lens[i]
	}
	best_tries := tries
	worst_tries := tries + lens[answer_len-1] - 1
	best := (k+5)*(best_tries/k) + (best_tries % k) + 1
	worst := (k+5)*(worst_tries/k) + (worst_tries % k) + 1

	fmt.Println(best, worst)
}
