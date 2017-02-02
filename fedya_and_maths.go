package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	divisor := 4
	re := regexp.MustCompile(`\r?\n`)

	reader := bufio.NewReader(os.Stdin)
	if _s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		s := re.ReplaceAllString(_s, "")
		remain := 0
		for i := 0; i < len(s); i++ {
			digit := int(s[i] - '0')
			dividend := digit + remain*10
			remain = dividend % divisor
		}

		if remain == 0 {
			fmt.Println(4)
		} else {
			fmt.Println(0)
		}
	}
}
