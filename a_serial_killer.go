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
	var n int
	var potential []string

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)

	if s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		t := strings.Split(re.ReplaceAllString(s, ""), " ")
		if len(t) != 2 {
			return
		}
		potential = t
	}

	if s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		u := re.ReplaceAllString(s, "")
		n, err = strconv.Atoi(u)
		if err != nil {
			return
		}
	}
	fmt.Printf("%s %s\n", potential[0], potential[1])

	for i := 0; i < n; i++ {
		if _in, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else {
			ab := strings.Split(re.ReplaceAllString(_in, ""), " ")
			if len(ab) != 2 {
				return
			}
			if potential[0] == ab[0] {
				potential[0] = ab[1]
			} else if potential[1] == ab[0] {
				potential[1] = ab[1]
			}
			fmt.Printf("%s %s\n", potential[0], potential[1])
		}
	}
}
