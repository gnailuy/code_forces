package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)

	if s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		s = re.ReplaceAllString(s, "")

		o := make([]rune, n)
		l := 0
		r := n - 1
		for i := 0; i < n; i++ {
			j := n - i - 1
			if i%2 == 0 {
				o[r] = rune(s[j])
				r--
			} else {
				o[l] = rune(s[j])
				l++
			}
		}

		fmt.Println(string(o))
	}
}
