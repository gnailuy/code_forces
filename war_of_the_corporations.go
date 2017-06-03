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

	var ai, phone string

	if s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		ai = re.ReplaceAllString(s, "")
	}
	if s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		phone = re.ReplaceAllString(s, "")
	}

	sharps := 0
	last_sharp_idx := -1

	pl := len(phone)
	for i := 0; i <= len(ai)-pl; i++ {
		sub := ai[i : i+pl]
		if sub == phone {
			if last_sharp_idx < 0 || i > last_sharp_idx {
				last_sharp_idx = i + pl - 1
				sharps += 1
			}
		}
	}

	fmt.Println(sharps)
}
