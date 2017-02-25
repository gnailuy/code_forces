package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	var k int
	if _, err := fmt.Scanf("%d\n", &k); err != nil {
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
		if k == 1 {
			fmt.Println(s)
			return
		}

		dict := make(map[byte]int)
		for i := 0; i < len(s); i++ {
			dict[s[i]] = dict[s[i]] + 1
		}

		var sub_string []byte

		for l, v := range dict {
			if v%k != 0 {
				fmt.Println(-1)
				return
			} else {
				for i := 0; i < v/k; i++ {
					sub_string = append(sub_string, l)
				}
			}
		}

		for i := 0; i < k; i++ {
			fmt.Printf("%s", string(sub_string))
		}
		fmt.Println()
	}
}
