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
		s := re.ReplaceAllString(s, "")
		m := len(s) / 2
		l := m
		r := m
		if len(s)%2 == 0 {
			l -= 1
		}

		diff := 0
		for i := 0; i <= l; i++ {
			if s[l-i] != s[r+i] {
				diff++
			}
		}
		if diff == 0 {
			if l == r {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		} else if diff == 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
