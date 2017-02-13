package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	var n, a, b int
	if _, err := fmt.Scanf("%d %d %d\n", &n, &a, &b); err != nil {
		return
	}

	re := regexp.MustCompile(`\r?\n`)

	reader := bufio.NewReader(os.Stdin)
	if s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		s = re.ReplaceAllString(s, "")
		if a == b || s[a-1] == s[b-1] {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
	}
}
