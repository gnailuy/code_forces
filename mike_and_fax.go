package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var s string
	var k int
	re := regexp.MustCompile(`\r?\n`)

	reader := bufio.NewReader(os.Stdin)
	if _s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		s = re.ReplaceAllString(_s, "")
	}

	if _k, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		_k = re.ReplaceAllString(_k, "")
		_vk, err := strconv.Atoi(_k)
		if err != nil {
			return
		}
		k = _vk
	}

	var ok bool = true
	l := len(s)
	if l%k != 0 {
		ok = false
	} else {
		n := l / k
		for i := 0; i < k; i++ {
			for j := 0; j < n/2; j++ {
				if s[n*i+j] != s[n*i+n-j-1] {
					ok = false
					break
				}
			}
			if !ok {
				break
			}
		}
	}

	if ok {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
