package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	reader := bufio.NewReader(os.Stdin)
	if s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		if n > 26 {
			fmt.Println(-1)
		} else if n < 2 {
			fmt.Println(0)
		} else {
			cnts := make([]int, 26)
			for i := 0; s[i] != '\r' && s[i] != '\n'; i++ {
				cnts[int(s[i]-'a')]++
			}
			dup_cnt := 0
			for i := 0; i < 26; i++ {
				if cnts[i] > 1 {
					dup_cnt += (cnts[i] - 1)
				}
			}
			fmt.Println(dup_cnt)
		}
	}
}
