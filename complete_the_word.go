package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	marks := make([]bool, 26)
	letter := 0
	question := 0
	ok := false

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)
	if _in, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		in := re.ReplaceAllString(_in, "")

		if len(in) < 26 {
			fmt.Println(-1)
			return
		}

		s := 0
		p := 0
		for s <= len(in)-25 {
			if in[s] == '?' {
				letter = 0
				question = 1
			} else {
				marks[in[s]-'A'] = true
				letter = 1
				question = 0
			}
			p = s + 1
			for p < len(in) && p-s < 26 {
				if in[p] == '?' {
					question++
				} else if marks[in[p]-'A'] {
					for i := 0; i < 26; i++ {
						marks[i] = false
					}
					j := p - 1
					for j > s && j < len(in) && (in[j] == '?' || in[j] != in[p]) {
						j--
					}
					s = j + 1
					break
				} else {
					marks[in[p]-'A'] = true
					letter++
				}
				p++
			}
			if letter+question == 26 {
				ok = true
				break
			}
		}
		if ok {
			i := 0
			for ; i < s; i++ {
				if in[i] == '?' {
					fmt.Printf("A")
				} else {
					fmt.Printf("%c", in[i])
				}
			}
			for ; i < p; i++ {
				if in[i] == '?' {
					for j := 0; j < 26; j++ {
						if !marks[j] {
							marks[j] = true
							fmt.Printf("%c", 'A'+j)
							break
						}
					}
				} else {
					fmt.Printf("%c", in[i])
				}
			}
			for ; i < len(in); i++ {
				if in[i] == '?' {
					fmt.Printf("A")
				} else {
					fmt.Printf("%c", in[i])
				}
			}
			fmt.Printf("\n")
		} else {
			fmt.Println(-1)
		}
	}
}
