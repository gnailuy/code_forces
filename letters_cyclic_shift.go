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

		start := false
		stop := false

		for i := 0; i < len(s); i++ {
			if !start && i == len(s)-1 {
				p := s[i] - 1
				if p < 'a' {
					p = 'z'
				}
				fmt.Printf("%c", p)
				break
			}
			if stop {
				fmt.Printf("%c", s[i])
				continue
			}
			p := s[i] - 1
			if p < 'a' {
				p = s[i]
				if start {
					stop = true
				}
			} else {
				start = true
			}
			fmt.Printf("%c", p)
		}
		fmt.Println()
	}
}
