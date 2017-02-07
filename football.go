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

	name := [2]string{"", ""}
	score := [2]int{0, 0}

	for i := 0; i < n; i++ {
		if _in, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else {
			in := re.ReplaceAllString(_in, "")
			if name[0] == "" || name[0] == in {
				name[0] = in
				score[0] += 1
			} else if name[1] == "" || name[1] == in {
				name[1] = in
				score[1] += 1
			}
		}
	}

	if score[0] > score[1] {
		fmt.Println(name[0])
	} else {
		fmt.Println(name[1])
	}
}
