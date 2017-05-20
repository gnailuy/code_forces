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

	var home string
	destinations := make(map[string]bool)

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)

	if s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		home = re.ReplaceAllString(s, "")
	}

	for i := 0; i < n; i++ {
		if s, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else {
			s = re.ReplaceAllString(s, "")
			ss := strings.Split(s, "->")
			var dest string
			if ss[0] == home {
				dest = ss[1]
			} else {
				dest = ss[0]
			}
			if destinations[dest] {
				delete(destinations, dest)
			} else {
				destinations[dest] = true
			}
		}
	}

	if len(destinations) == 0 {
		fmt.Println("home")
	} else {
		fmt.Println("contest")
	}
}
