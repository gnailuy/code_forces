package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	const mmd string = "Bulbasaur"
	var n int
	var ss []byte

	re := regexp.MustCompile(`\r?\n`)

	reader := bufio.NewReader(os.Stdin)
	if _s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		_s = re.ReplaceAllString(_s, "")
		n = len(_s)
		ss = make([]byte, n)
		for i := 0; i < n; i++ {
			ss[i] = _s[i]
		}
	}

	var count int = 0
	var p_mmd int = 0
	var p_s int = 0
	var query_cnt int = 0
	for true {
		if mmd[p_mmd] == ss[p_s] {
			ss[p_s] = ' '
			query_cnt = 0
			p_mmd++
			if p_mmd == len(mmd) {
				p_mmd = 0
				count++
			}
		} else {
			p_s++
			if p_s == n {
				p_s = 0
			}
			query_cnt++
			if query_cnt > n {
				break
			}
		}
	}

	fmt.Println(count)
}
