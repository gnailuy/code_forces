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

type Entry struct {
	name  string
	score int
}

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	history := make([]Entry, n)
	scoreboard := make(map[string]int)

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)
	for i := 0; i < n; i++ {
		if s, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else {
			s = re.ReplaceAllString(s, "")

			ss := strings.Split(s, " ")
			name := ss[0]
			score, err := strconv.Atoi(ss[1])
			if err != nil {
				return
			}

			history = append(history, Entry{name, score})
			scoreboard[name] = scoreboard[name] + score
		}
	}

	max := 0
	for _, score := range scoreboard {
		if score > max {
			max = score
		}
	}
	winners := make(map[string]int)
	for name, score := range scoreboard {
		if score == max {
			winners[name] = score
		}
	}

	scoreboard = make(map[string]int)
	for _, entry := range history {
		if winners[entry.name] > 0 {
			scoreboard[entry.name] = scoreboard[entry.name] + entry.score
			if scoreboard[entry.name] >= max {
				fmt.Println(entry.name)
				break
			}
		}
	}
}
