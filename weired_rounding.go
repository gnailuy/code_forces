package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)

	if s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		nk := strings.Split(re.ReplaceAllString(s, ""), " ")
		n := nk[0]
		l := len(n)
		k, err := strconv.Atoi(nk[1])
		if err != nil {
			return
		}

		remove := 0
		cnt_0 := 0
		for i := l - 1; i >= 0; i-- {
			if n[i] == '0' {
				cnt_0++
				if cnt_0 == k {
					break
				}
			} else {
				remove++
			}
		}

		if cnt_0 < k {
			fmt.Println(l - 1)
		} else {
			fmt.Println(remove)
		}
	}
}
