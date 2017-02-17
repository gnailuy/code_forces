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

		keys := map[byte]int{}
		buy_key := 0
		for i := 0; i < n-1; i++ {
			key := s[i*2]
			lock := s[i*2+1] + ('a' - 'A')
			if key == lock {
				continue
			} else {
				if _, exist := keys[lock]; exist && keys[lock] > 0 {
					keys[lock] = keys[lock] - 1
				} else {
					buy_key++
				}
				if _, exist := keys[key]; exist {
					keys[key] = keys[key] + 1
				} else {
					keys[key] = 1
				}
			}
		}

		fmt.Println(buy_key)
	}
}
