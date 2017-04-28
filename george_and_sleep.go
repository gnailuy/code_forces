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

	var st [2][2]int

	for i := 0; i < 2; i++ {
		if _in, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else {
			hm := strings.Split(re.ReplaceAllString(_in, ""), ":")
			var err error
			for j := 0; j < 2; j++ {
				st[i][j], err = strconv.Atoi(hm[j])
				if err != nil {
					return
				}
			}
		}
	}

	h := st[0][0] - st[1][0]
	m := st[0][1] - st[1][1]
	if m < 0 {
		m += 60
		h -= 1
	}
	if h < 0 {
		h += 24
	}

	fmt.Printf("%02d:%02d\n", h, m)
}
