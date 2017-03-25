package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)

	if _in, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		word := re.ReplaceAllString(_in, "")
		set := make(map[string]bool)

		n := len(word)
		for i := 0; i < n; i++ {
			new := strings.Join([]string{word[n-i:], word[:n-i]}, "")
			set[new] = true
		}
		fmt.Println(len(set))
	}
}
