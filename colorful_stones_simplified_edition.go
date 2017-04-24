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

	var s, t string

	if in, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		s = re.ReplaceAllString(in, "")
	}

	if in, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		t = re.ReplaceAllString(in, "")
	}

	pos := 0
	for i := 0; i < len(t); i++ {
		if t[i] == s[pos] {
			pos++
		}
	}

	fmt.Println(pos + 1)
}
