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

		letter_count_map := make(map[rune]int)
		for i := 0; i < len(s); i++ {
			letter_count_map[rune(s[i])] += 1
		}

		odd_count := 0
		for _, c := range letter_count_map {
			if c%2 != 0 {
				odd_count += 1
			}
		}

		if odd_count == 0 || odd_count%2 != 0 {
			fmt.Println("First")
		} else {
			fmt.Println("Second")
		}
	}
}
