package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	ab := [2]string{}

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)
	for i := 0; i < 2; i++ {
		if s, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else {
			ab[i] = re.ReplaceAllString(s, "")
		}
	}

	if len(ab[0]) > len(ab[1]) {
		fmt.Println(len(ab[0]))
	} else if len(ab[0]) < len(ab[1]) {
		fmt.Println(len(ab[1]))
	} else {
		if ab[0] == ab[1] {
			fmt.Println(-1)
		} else {
			fmt.Println(len(ab[0]))
		}
	}
}
