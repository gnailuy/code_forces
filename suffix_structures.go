package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)

	st := [2]string{}

	for i := 0; i < 2; i++ {
		if _in, err := reader.ReadString('\n'); err != nil {
			if err != io.EOF {
				return
			}
		} else {
			st[i] = re.ReplaceAllString(_in, "")
		}
	}

	s := st[0]
	t := st[1]

	if len(s) < len(t) {
		fmt.Println("need tree")
	} else {
		j := 0
		auto := true
		for i := 0; i < len(t); i++ {
			if !auto {
				break
			}
			for true {
				if j >= len(s) {
					auto = false
					break
				} else if s[j] == t[i] {
					j++
					break
				}
				j++
			}
		}
		if auto {
			fmt.Println("automaton")
			return
		}

		sm := make(map[rune]int)
		for i := 0; i < len(s); i++ {
			sm[rune(s[i])] += 1
		}
		for j := 0; j < len(t); j++ {
			sm[rune(t[j])] -= 1
		}

		remain := 0
		has_neg := false
		for _, v := range sm {
			if v < 0 {
				has_neg = true
			}
			remain += v
		}
		if has_neg {
			fmt.Println("need tree")
		} else if remain == 0 {
			fmt.Println("array")
		} else {
			fmt.Println("both")
		}
	}
}
