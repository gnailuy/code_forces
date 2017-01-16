package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	var longest int = 0
	var words_inside int = 0

	reader := bufio.NewReader(os.Stdin)
	if s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		inside := false
		word_length := 0
		for i := 0; i < n; i++ {
			if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
				word_length++
			} else {
				if !inside && word_length > longest {
					longest = word_length
				}
				if inside && word_length > 0 {
					words_inside++
				}
				if s[i] == '(' {
					inside = true
				} else if s[i] == ')' {
					inside = false
				}
				word_length = 0
			}
		}
		if word_length > longest {
			longest = word_length
		}
	}

	fmt.Printf("%d %d\n", longest, words_inside)
}
