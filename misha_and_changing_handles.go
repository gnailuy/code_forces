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
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	used := make(map[string]bool)
	new_old := make(map[string]string)

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)
	for i := 0; i < n; i++ {
		if s, err := reader.ReadString('\n'); err != nil {
			if err != io.EOF {
				return
			}
		} else {
			ss := strings.Split(re.ReplaceAllString(s, ""), " ")
			if len(ss) != 2 {
				return
			}
			oh := ss[0]
			nh := ss[1]

			used[oh] = true
			if !used[nh] {
				if ooh, ok := new_old[oh]; ok {
					new_old[nh] = ooh
					delete(new_old, oh)
				} else {
					new_old[nh] = oh
				}
				used[nh] = true
			}
		}
	}

	fmt.Println(len(new_old))
	for nh, oh := range new_old {
		fmt.Println(oh, nh)
	}
}
