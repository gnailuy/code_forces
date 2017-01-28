package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`\r?\n`)
	strs := make([]string, 2)

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 2; i++ {
		if s, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else {
			strs[i] = re.ReplaceAllString(s, "")
		}
	}
	out_str := make([]byte, len(strs[0]))

	follow := 0
	for i := 0; i < len(strs[0]); i++ {
		if strs[0][i] == strs[1][i] {
			out_str[i] = strs[0][i]
		} else {
			out_str[i] = strs[follow][i]
			follow = (follow + 1) % 2
		}
	}

	if follow == 1 {
		fmt.Println("impossible")
	} else {
		fmt.Printf("%s\n", out_str)
	}
}
